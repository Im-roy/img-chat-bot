package main

import (
	"img-chat-bot/config"
	"img-chat-bot/db"
	"img-chat-bot/server"
	"img-chat-bot/utils"
)

func main() {
	logger := utils.NewLogObject()
	logger.Info("img-chat-bot app started...")

	appConfig, err := config.GetConfig("img-chat-bot-config.yml")
	if err != nil {
		logger.Error("error in building app config")
		panic(err)
	}

	dbClient, err := db.GetDBGormClient(appConfig.DbConfig)
	if err != nil {
		logger.Error("error in creating db client")
		panic(err)
	}

	// start http server
	httpHandler := server.HttpHandler{}
	httpHandler.Init(dbClient)
}
