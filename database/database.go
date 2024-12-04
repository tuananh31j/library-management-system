package database

import (
	"fmt"
	"time"

	"github.com/tuananh31j/library-management-system/config"
	"github.com/tuananh31j/library-management-system/model"
	"github.com/tuananh31j/library-management-system/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log.Fatalf("Failed to connect to database: %+v", err)
	}

	// Config pooling
	sqlDB, errDB := dbInstance.DB()
	if errDB != nil {
		utils.Log.Fatalf("Failed to get db instance: %+v", errDB)
	}
	dbInstance.AutoMigrate(&model.Author{})
	dbInstance.AutoMigrate(&model.Book{})
	dbInstance.AutoMigrate(&model.Borrower{})
	dbInstance.AutoMigrate(&model.BorrowerBooks{})

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	return dbInstance
}
