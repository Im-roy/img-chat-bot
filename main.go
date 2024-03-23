package main

import (
	"fmt"
	"img-chat-bot/db"
	"img-chat-bot/server"
)

func main() {
	fmt.Println("app started...")

	dbClient, err := db.GetDBGormClient()
	if err != nil {
		panic(err)
	}
	fmt.Println(dbClient)
	// start http server
	httpHandler := server.HttpHandler{}
	httpHandler.Init()
}
