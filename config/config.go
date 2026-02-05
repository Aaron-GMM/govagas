package config

import "gorm.io/gorm"

// variavel do package
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	return nil
}

func GetLogger(p string) *Logger {
	// initialize logger
	logger = NewLogger(p)
	return logger
}
