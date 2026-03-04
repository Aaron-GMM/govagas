package main

import (
	"github.com/Aaron-GMM/govagas/config"
	"github.com/Aaron-GMM/govagas/router"
	"github.com/joho/godotenv"
)

var (
	logger config.Logger
)

func main() {
	_ = godotenv.Load()
	logger = *config.GetLogger("main")

	var err error
	err = config.Init()

	if err != nil {
		logger.ErrorF("config init err: %v", err)
		return
	}

	// incializando router
	router.InitRouter()
}
