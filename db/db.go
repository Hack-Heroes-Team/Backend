package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	port := os.Getenv("DB_CRED")
	db, err := gorm.Open(postgres.Open(port), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10000000)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
