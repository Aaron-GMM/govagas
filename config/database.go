package config

import (
	"os"

	"github.com/Aaron-GMM/govagas/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	//checke if the database file exist and Create DataBase
	// path for db
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		logger.ErrorF("Can't find DB_PATH environment variable: %v", dbPath)
	}
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Data base file not found, Creating new database")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		file, err := os.Create(dbPath)

		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		file.Close()
	}

	//create db and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error opening database: %v", err)
		return nil, err
	}

	// Migrates the Schemas
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("Error auto-migrating database: %v", err)
	}
	return db, nil
}
