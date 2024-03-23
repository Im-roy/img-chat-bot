package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HttpRoutesHandler struct {
	Router *mux.Router
}

func (h *HttpRoutesHandler) RegisterRoutes() {
	h.Router.HandleFunc("/ping", h.HandlePing).Methods(http.MethodGet)
	h.Router.HandleFunc("/add-images/v1", h.HandleAddImages).Methods(http.MethodPost)
	h.Router.HandleFunc("/get-images/v1", h.HandleGetImages).Methods(http.MethodGet)
}
