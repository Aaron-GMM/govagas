package router

import (
	"net/http"

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
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": " GET opening",
			})
		})
		// create Oportunity
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": " CREATE opening",
			})
		})

		// update Oportunity
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": " Update opening",
			})
		})

		// delete Oportunity
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": " DELETE opening",
			})
		})

		// show all Oportunity
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": " GET openings",
			})
		})
	}
}
