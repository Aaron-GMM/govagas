package router

import "github.com/gin-gonic/gin"

func InitRouter() {

	// initialize my gin
	var router *gin.Engine = gin.Default()

	// my routes
	initializeRouter(router)
	// run application
	router.Run(":8080")
}
