package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
	"time"
	"zaimik/internal/app/models"
	"zaimik/internal/app/service"
	"zaimik/internal/pkg/smtp"
)

const (
	adminAuthCodeTmplName = "admin_auth_code.html"
	adminCookiePath       = "/admin/"
	adminSessLifeDuration = 60 * 60 * 24
	adminSessCookieName   = "admin_session_id"
)

type adminAuthReq struct {
	Login    string `json:"admin_login" binding:"required"`
	Password string `json:"admin_password" binding:"required"`
}

func (h *Handler) checkAdmin(ctx *gin.Context) {
	req := adminAuthReq{}
	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Errorf("error while binding json in admin API: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "bad format of request")
		return
	}

	if req.Login != os.Getenv("ADMIN_LOGIN") || req.Password != os.Getenv("ADMIN_PASSWORD") {
		h.logger.Errorf("somebody want to log in as admin with incorrect params")
		newRespErr(ctx, http.StatusForbidden, "forbidden")
		return
	}

	_, err := h.services.Authorization.CheckEmailAndSendAuthCode(os.Getenv("ADMIN_MAIL"), adminAuthCodeTmplName, "Вход в админ панель на сайте zaimik")
	switch err {
	case smtp.BadEmail:
		h.logger.Infof("user sent incorect email: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "error while sending mail")
		return
	case nil:
		ctx.JSON(http.StatusOK, gin.H{})
	default:
		h.logger.Errorf("error sending mail: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "error while sending mail")
		return
	}
}

func (h *Handler) authAdmin(ctx *gin.Context) {
	req := struct {
		Code string `json:"code" binding:"required"`
	}{}
	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Infof("error binding json: %s", err.Error())
		newRespErr(ctx, http.StatusForbidden, "incorrect auth code")
		return
	}

	sid, err := h.services.Administration.AuthorizeAdmin(req.Code)
	switch err {
	case service.IncorrectAuthCode:
		newRespErr(ctx, http.StatusBadRequest, "auth code had been expired or incorrect")
		return
	case nil:
		ctx.SetCookie(adminSessCookieName, sid, adminSessLifeDuration, adminCookiePath, viper.GetString("http.domain"), true, true)
		ctx.Redirect(http.StatusSeeOther, "/admin/panel/")
		return
	default:
		h.logger.Errorf("error while trying authorize admin: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (h *Handler) signOutAdmin(ctx *gin.Context) {
	//_, ok := ctx.Get(adminCtx)
	//if !ok {
	//	h.logger.Infof("someone want to signOutAdmin without adminCtx")
	//	newRespErr(ctx, http.StatusForbidden, "unauthorized")
	//	return
	//}

	sid, err := ctx.Cookie(adminSessCookieName)
	if err != nil {
		h.logger.Errorf("admin without session: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "have no session")
		return
	}

	if err := h.services.Administration.LogOutAdmin(sid); err != nil {
		h.logger.Errorf("error while logout admin: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	ctx.SetCookie(adminSessCookieName, "", -1, adminCookiePath, viper.GetString("http.domain"), true, true)
	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin", viper.GetString("http.protocol")))
}

func (h *Handler) uploadCompany(ctx *gin.Context) {
	var company models.LoanCompanyAdmin
	if err := ctx.Bind(&company); err != nil {
		h.logger.Errorf("error binding formdata: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "bad formdata")
		return
	}

	logo := ctx.Request.MultipartForm.File["logo"][0]
	company.LogoNameOnS3 = logo.Filename

	logoReader, err := logo.Open()
	company.Logo = logoReader

	if err != nil {
		h.logger.Errorf("error getting logo from formdata")
		newRespErr(ctx, http.StatusBadRequest, "bad logo formatting")
		return
	}

	if _, err := h.services.Administration.AddCompany(company); err != nil {
		h.logger.Errorf("error while adding company: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "invalid fields")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/companies/", viper.GetString("http.protocol")))
}

func (h *Handler) updateCompanies(ctx *gin.Context) {
	var req []models.LoanCompanyPriorityUpdate
	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Errorf("error binding json into []models.LoanCompanyPriorityUpdate: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "bad json")
		return
	}

	if err := h.services.Administration.UpdateCompaniesPriority(req); err != nil {
		h.logger.Errorf("error updating priority: %s", err)
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/companies/", viper.GetString("http.protocol")))
}

func (h *Handler) refactorCompany(ctx *gin.Context) {
	var company models.LoanCompanyAdmin
	if err := ctx.BindJSON(&company); err != nil {
		h.logger.Errorf("error while binding json: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "bad json")
		return
	}

	id := ctx.Param("id")
	companyId, err := strconv.Atoi(id)
	if err != nil || companyId <= 0 {
		newRespErr(ctx, http.StatusBadRequest, "bad url params")
		return
	}

	company.Id = companyId

	if err := h.services.Administration.RefactorCompany(company); err != nil {
		h.logger.Errorf("error while refactoring company with id = %d: %s", company.Id, err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/companies/", viper.GetString("http.protocol")))
}

func (h *Handler) deleteCompany(ctx *gin.Context) {
	id := ctx.Param("id")
	companyId, err := strconv.Atoi(id)
	if err != nil || companyId <= 0 {
		newRespErr(ctx, http.StatusBadRequest, "bad params")
		return
	}

	if err := h.services.Administration.DeleteCompanyById(ctx, companyId); err != nil {
		h.logger.Errorf("error while deleting company with id = %d: %s", companyId, err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/companies/", viper.GetString("http.protocol")))
}

func (h *Handler) updateReview(ctx *gin.Context) {
	var review models.ReviewAdmin
	if err := ctx.BindJSON(&review); err != nil {
		h.logger.Errorf("error binding review: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "incorrect request format")
		return
	}

	if _, err := h.services.Administration.SetReviewModerated(review); err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusBadRequest, "invalid json")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/reviews/", viper.GetString("http.protocol")))
}

func (h *Handler) uploadReview(ctx *gin.Context) {
	var review models.ReviewAdmin
	if err := ctx.BindJSON(&review); err != nil {
		h.logger.Errorf("error binding review: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "incorrect request format")
		return
	}

	if _, err := h.services.Administration.UploadReview(review); err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusBadRequest, "invalid json")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/reviews/", viper.GetString("http.protocol")))
}

type deleteReviewReq struct {
	Id int `json:"id" binding:"required" db:"id"`
}

func (h *Handler) deleteReview(ctx *gin.Context) {
	var req deleteReviewReq
	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Errorf("error binding json into deleteReviewReq: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "invalid json")
		return
	}

	if err := h.services.Administration.DeleteReview(req.Id); err != nil {
		h.logger.Error(err)
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/reviews/", viper.GetString("http.protocol")))
}

type refactorSubscriptionReq struct {
	Id          int    `json:"id" binding:"required"`
	ExpiredDate string `json:"expired_date" binding:"required"`
}

func (h *Handler) refactorSubscription(ctx *gin.Context) {
	var sub refactorSubscriptionReq
	if err := ctx.BindJSON(&sub); err != nil {
		h.logger.Errorf("error binding json: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "invalid json")
		return
	}

	expiredDate, err := time.Parse(time.RFC3339, sub.ExpiredDate)
	if err != nil {
		h.logger.Errorf("error parsing time: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "invalid json")
		return
	}

	err = h.services.LoanCompaniesManager.RefactorSubscriptionExpiredDate(sub.Id, &expiredDate)
	if err != nil {
		h.logger.Errorf("error while updating subscription`s expired_date with id = %d: %s", sub.Id, err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/admin/panel/subscriptions/", viper.GetString("http.protocol")))
}
