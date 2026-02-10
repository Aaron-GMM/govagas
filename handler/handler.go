package handler

import (
	"github.com/Aaron-GMM/govagas/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	Db = config.GetDB()
}
