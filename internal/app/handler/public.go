package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"zaimik/internal/app/models"
	"zaimik/internal/app/service"
	"zaimik/internal/pkg/smtp"
)

const (
	sessionId                    = "session_id"
	sessionLifeDuration          = 60 * 60 * 24 * 14
	emailTemplateForSendAuthCode = "test.html"
)

type getCodeReq struct {
	Email string `json:"email" binding:"required"`
}

// @Summary getAuthCode
// @Tags publicAuth
// @Description api endpoint that receives the user's mail, an authorization code is sent to it, which must be entered in a special field, the code lives for 1 minute. responds with status 400 if invalid data is sent or invalid mail is sent
// @ID
// @Accept json
// @Produce json
// @Param input body getCodeReq true "email"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /get-code [post]
func (h *Handler) getAuthCode(ctx *gin.Context) {
	var req getCodeReq

	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Errorf("error whihle binding json: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "incorrect request format")
		return
	}

	_, err := h.services.Authorization.CheckEmailAndSendAuthCode(req.Email, emailTemplateForSendAuthCode, "Закончите регистрацию на сайте zaimik")
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

type signInReq struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

// @Summary signIn
// @Tags publicAuth
// @Description api endpoint that accepts the user's mail along with the authorization code, checks the correctness of the code and sets cookies (session_id) in the browser. responds with status 200 if successful, with status 400 if code is invalid or email is invalid
// @ID
// @Accept json
// @Produce json
// @Param input body signInReq true "email and code"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router / [post]
func (h *Handler) signIn(ctx *gin.Context) {
	var req signInReq

	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Errorf("error whihle binding json: %s", err.Error())
		newRespErr(ctx, http.StatusBadRequest, "incorrect request fromat")
		return
	}

	sid, err := h.services.Authorization.AuthorizeUser(req.Email, req.Code)
	switch err {
	case service.IncorrectAuthCode:
		h.logger.Infof("user sent incorrect auth code")
		newRespErr(ctx, http.StatusBadRequest, "auth code had been expired or incorrect")
		return
	case nil:
		ctx.SetCookie(sessionId, sid, sessionLifeDuration, "/", viper.GetString("http.domain"), true, false)
		ctx.JSON(http.StatusOK, gin.H{})
		return
	default:
		h.logger.Errorf("unknown error while authorizing user: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
}

type reviewsResponse struct {
	Reviews []models.Review `json:"reviews"`
}

// @Summary getReviews
// @Tags public
// @Description api endpoint, send all moderated reviews which consist in dataBase in json format
// @ID
// @Produce json
// @Success 200 {array} reviewsResponse
// @Failure 500 {object} errorResponse
// @Router /reviews [get]
func (h *Handler) getReviews(ctx *gin.Context) {
	reviews, err := h.services.GetReviews()
	if err != nil {
		h.logger.Errorf("error while getting reviews: %s", err.Error())
		newRespErr(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, reviewsResponse{Reviews: reviews})
}
