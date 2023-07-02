package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) indexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "auth.html", gin.H{})
}
