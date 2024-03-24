package db

import (
	"fmt"
	"img-chat-bot/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBGormClient(dbConfig config.DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
