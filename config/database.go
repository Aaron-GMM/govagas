package config

import (
	"fmt"
	"os"

	"github.com/Aaron-GMM/govagas/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("postegres")
	//checke if the database file exist and Create DataBase
	// path for db
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || user == "" || dbname == "" {
		logger.ErrorF("Missing database environment variables")
		return nil, fmt.Errorf("missing database environment variables")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)
	logger.Info("Connecting to Postgres database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error connecting to database: %v", err)
		return nil, err
	}
	logger.Info("Connected successfully!")
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("Error auto-migrating database: %v", err)
		return nil, err
	}

	logger.Info("Migrations executed successfully!")
	return db, nil
}
