package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:pass@localhost:5432/hack"), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10000000)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
