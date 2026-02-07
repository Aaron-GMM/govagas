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
	if err := godotenv.Load(); err != nil {
		// Tratar erro ou apenas logar
		config.GetLogger("main").Info("No .env file found or error loading it")
	}
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
