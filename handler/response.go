package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "aplication/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"ErrorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data":    data,
	})
}
