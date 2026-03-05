package handler

import (
	"github.com/Aaron-GMM/govagas/config"
)

var (
	Logger *config.Logger
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
}
