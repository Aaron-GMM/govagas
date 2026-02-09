package openingHandler

import (
	"github.com/Aaron-GMM/govagas/config"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := struct {
		role string
	}{}

	ctx.BindJSON(&request)

	config.Logger.InforF("request received: %+v", request)

}
