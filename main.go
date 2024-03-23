package main

import (
	"fmt"
	"img-chat-bot/server"
)

func main() {
	fmt.Println("app started...")

	// start http server
	httpHandler := server.HttpHandler{}
	httpHandler.Init()
}
