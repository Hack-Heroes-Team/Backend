package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbCred := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbCred), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10000000)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
