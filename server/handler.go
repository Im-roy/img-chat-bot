package server

import (
	aimodel "img-chat-bot/AIModel"
	"img-chat-bot/AIModel/gemini"
	"img-chat-bot/chatbot"
	dbRepo "img-chat-bot/repo/dbRepo"
	fileRepo "img-chat-bot/repo/fileRepo"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type HttpHandler struct {
	handler http.Handler
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) Init(db *gorm.DB) {
	apiRouter := mux.NewRouter()
	http.Handle("/", apiRouter)

	httpRouter := HttpRoutesHandler{
		Router: apiRouter,
		ChatBot: chatbot.ChatBot{
			AIModel:  aimodel.AiModel{AIClient: gemini.GeminiAI{}},
			FileRepo: fileRepo.FileRepo{},
			DbRepo: dbRepo.DbRepo{
				DB: db,
			},
		},
	}
	httpRouter.RegisterRoutes()
	h.ServeHttp()
}

func (h *HttpHandler) ServeHttp() {
	log.Println("started api server...")
	srv := &http.Server{
		Handler:      h.handler,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error in calling listen and serve")
	}
}
