package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBGormClient() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=postgres dbname=chat-bot port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
