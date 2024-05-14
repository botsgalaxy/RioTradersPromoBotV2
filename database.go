package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PromoBotUser struct {
	gorm.Model
	UserId       int64 `gorm:"unique"`
	FirstName    string
	LastName     string
	Username     string
	IsPremium    bool
	LanguageCode string
}

var DB *gorm.DB

func migrateDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("<< Failed to connect database . Exiting Now >>")
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = DB.AutoMigrate(&PromoBotUser{})
	if err != nil {
		panic("<< Failed to Automigrate Models >>")
	}

}
