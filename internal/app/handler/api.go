package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/app/service"
)

func (h *Handler) chatPage(ctx *gin.Context) {
	if _, ok := ctx.Get(userCtx); !ok {
		newRespErr(ctx, http.StatusUnauthorized, "not authorized")
		return
	}
	ctx.HTML(http.StatusOK, "chatPage.html", gin.H{})
}

type updateReq struct {
	User models.User `json:"user"`
	Loan models.Loan `json:"loan"`
}

type suitableCompanies struct {
	Companies []models.LoanCompany `json:"companies"`
}

// @Summary updateUserGetLinkOnPayment
// @Tags unfinished
// @Description unfinished, because the online cash register is not connected, do not use this
// @ID
// @Accept json
// @Produce json
// @Param input body updateReq true "user update"
// @Success 200 {array} suitableCompanies
// @Success 204 {array} suitableCompanies
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/catalog [put]
func (h *Handler) updateUserGetLinkOnPayment(ctx *gin.Context) {
	userId, ok := ctx.Get(userCtx)
	if !ok {
		newRespErr(ctx, http.StatusUnauthorized, "not authorized")
		return
	}

	var req updateReq

	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Errorf("error while binding json: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "incorrect format of request")
		return
	}

	uid, ok := userId.(int)
	if !ok || uid <= 0 {
		newRespErr(ctx, http.StatusBadRequest, "bad uid")
		return
	}

	req.Loan.UserId = uid
	req.User.Id = uid

	loanId, err := h.services.LoanCompaniesManager.AddLoanRequest(req.Loan)
	if err != nil {
		h.logger.Errorf("error creating loan request for userId = %d: %s", uid, err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	subscription, err := h.services.LoanCompaniesManager.InitSubscription(uid, loanId)
	if err != nil {
		h.logger.Errorf("error while init subscription: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	fmt.Println(subscription)

	if err := h.services.Authorization.UpdateUser(req.User); err != nil {
		h.logger.Errorf("user update was failed: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	// TODO здесь мы должны отдать ссылку на оплату и если успешно, то записать в репозиторий обновленную подписку
}

// @Summary getSortedSuitableCatalog
// @Tags api
// @Description when user successfully pay for subscription, he can get sorted catalog of loan companies
// @ID
// @Accept json
// @Produce json
// @Param input body updateReq true "user update"
// @Success 200 {array} suitableCompanies
// @Success 204 {array} suitableCompanies
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/catalog [post]
func (h *Handler) getSortedSuitableCatalog(ctx *gin.Context) {
	userId, ok := ctx.Get(userCtx)
	if !ok {
		newRespErr(ctx, http.StatusUnauthorized, "not authorized")
		return
	}

	id, ok := userId.(int)
	if !ok || id <= 0 {
		newRespErr(ctx, http.StatusBadRequest, "bad id")
		return
	}

	companies, err := h.services.LoanCompaniesManager.GetSortedSuitableCatalogByUserId(id)
	if err == nil {
		ctx.JSON(http.StatusOK, suitableCompanies{Companies: companies})
	} else if err == service.HaveNoSuitableCompanies {
		ctx.JSON(http.StatusNoContent, suitableCompanies{Companies: companies})
	} else {
		h.logger.Errorf("unknown error while getting catalog: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "error while getting catalog")
		return
	}
}

type checkSubscriptionResponse struct {
	Active      bool      `json:"active"`
	ExpiredDate time.Time `json:"expired_date"`
}

// @Summary checkSubscription
// @Tags api
// @Description after payment for subscription endpoint can answer what status of subscription is it now
// @ID
// @Accept json
// @Produce json
// @Success 200 {object} checkSubscriptionResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/check-subscription [get]
func (h *Handler) checkSubscription(ctx *gin.Context) {
	userId, ok := ctx.Get(userCtx)
	if !ok {
		newRespErr(ctx, http.StatusUnauthorized, "not authorized")
		return
	}

	id, ok := userId.(int)
	if !ok || id <= 0 {
		newRespErr(ctx, http.StatusBadRequest, "bad id")
		return
	}

	active, expiredDate, err := h.services.LoanCompaniesManager.CheckSubscriptionByUserId(id)
	if err != nil && err != service.InvalidModel && err != service.HaveNoSubscription {
		h.logger.Errorf("unknown error occured: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	} else if err == service.InvalidModel || err == service.HaveNoSubscription {
		newRespErr(ctx, http.StatusBadRequest, "invalid data")
		return
	} else {
		ctx.JSON(http.StatusOK, checkSubscriptionResponse{Active: active, ExpiredDate: expiredDate})
	}
}

type sendReviewRequest struct {
	Review models.Review `json:"review" binding:"required"`
}

// @Summary sendReview
// @Tags api
// @Description api endpoint, accepts feedback from a person if he is authorized, if not authorized, responds with an error
// @ID
// @Accept json
// @Produce json
// @Param input body sendReviewRequest true "info of review"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/send-review [post]
func (h *Handler) sendReview(ctx *gin.Context) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		h.logger.Errorf("user want to send request without session")
		newRespErr(ctx, http.StatusUnauthorized, "have no session")
		return
	}

	var req sendReviewRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Errorf("error while binding json from user_id = %s: %s", id, err.Error())
		newRespErr(ctx, http.StatusBadRequest, "bad request")
		return
	}

	if err := h.services.LoanCompaniesManager.AddReview(req.Review); err != nil {
		h.logger.Errorf("error adding review from user_id = %s: %s", id, err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "error adding review")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// @Summary signOut
// @Tags api
// @Description api endpoint, logs out the user if he is logged in
// @ID
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 500 {object} errorResponse
// @Router /api/sign-out [post]
func (h *Handler) signOut(ctx *gin.Context) {
	sid, err := ctx.Cookie("session_id")
	if err != nil {
		h.logger.Infof("user want to sign-out without session: %s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	userId, ok := ctx.Get(userCtx)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{})
	}

	numId, ok := userId.(int)
	if !ok {
		h.logger.Errorf("can`t cast userCtx to int")
		newRespErr(ctx, http.StatusInternalServerError, "bad id")
		return
	}

	if err := h.services.Authorization.DeleteSession(sid, numId); err != nil {
		h.logger.Errorf("error while deleting session, responsed with status 500: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "error while signing out")
		return
	}

	ctx.SetCookie(sessionId, sid, -1, "/", viper.GetString("http.domain"), true, false)
	ctx.JSON(http.StatusOK, gin.H{})
}
