package router

import (
	"github.com/Aaron-GMM/govagas/handler"
	"github.com/gin-gonic/gin"
)

/*
 point utilizado (router *gin.Engine) em vez de ficar passando o router diretamente repassso apenas o poin
que ele se localiza para acessar ele sendo bem mais eficiente
em vez de acessar uma instacia do router para cada função acesso apenas um router sendo bem mais eficiente
*/

func initializeRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// show Oportunity
		v1.GET("/opening", handler.ShowOpeningHandler)
		// create Oportunity
		v1.POST("/opening", handler.CreateOpeningHandler)

		// update Oportunity
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// delete Oportunity
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		// show all Oportunity
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
