package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Msg string `json:"msg"`
}

func newRespErr(ctx *gin.Context, statusCode int, msg string) {
	ctx.AbortWithStatusJSON(statusCode, errorResponse{Msg: msg})
}
