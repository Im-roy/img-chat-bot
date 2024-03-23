package server

import "net/http"

func (h *HttpRoutesHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "server is up and running"}`))
}

func (h *HttpRoutesHandler) HandleAddImages(w http.ResponseWriter, r *http.Request) {
	panic("to be implemented...")
}

func (h *HttpRoutesHandler) HandleGetImages(w http.ResponseWriter, r *http.Request) {
	panic("to be implemented...")
}
