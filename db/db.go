package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	db, err := gorm.Open(postgres.Open("postgres://daamrfehofxpau:1dd396b925c7b2655ec3dc8abb72a016cc1f9bf3e38260b637f89ba6b8635e74@ec2-34-248-169-69.eu-west-1.compute.amazonaws.com:5432/delp14k9rbbup"), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10000000)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
