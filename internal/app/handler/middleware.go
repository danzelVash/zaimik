package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userCtx  = "user_id"
	adminCtx = "admin_id"
)

func (h *Handler) checkSession(ctx *gin.Context) {
	if sid, err := ctx.Cookie(sessionId); err == nil {
		if id, err := h.services.Authorization.CheckSession(sid); err == nil {
			ctx.Set(userCtx, id)
			return
		}
	}
	//newRespErr(ctx, http.StatusUnauthorized, "session had been expired or incorrect")
}

func (h *Handler) checkAdminSession(ctx *gin.Context) {
	if adminSid, err := ctx.Cookie(adminSessCookieName); err == nil {
		if id, err := h.services.Administration.CheckAdminSession(adminSid); err == nil {
			ctx.Set(adminCtx, id)
			return
		} else {
			h.logger.Infof("someone want to log in as admin without admin session cookie")
			newRespErr(ctx, http.StatusForbidden, "no information on this session")
			return
		}
	} else {
		h.logger.Infof("someone want to log in as admin without admin session cookie")
		newRespErr(ctx, http.StatusForbidden, "no information on this session")
		return
	}
}
