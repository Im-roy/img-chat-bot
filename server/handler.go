package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HttpHandler struct {
	handler http.Handler
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) Init() {
	apiRouter := mux.NewRouter()
	http.Handle("/", apiRouter)

	httpRouter := HttpRoutesHandler{
		Router: apiRouter,
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
