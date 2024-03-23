package server

import (
	"encoding/json"
	"img-chat-bot/model"
	"io/ioutil"
	"net/http"
)

func (h *HttpRoutesHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "server is up and running"}`))
}

func (h *HttpRoutesHandler) HandleAddImages(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	requestBody := model.RequestModel{}
	err = json.Unmarshal(reqBody, &requestBody)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	resp, err := h.ChatBot.GenerateResponse(ctx, requestBody.Prompt, 1)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
	return
}

func (h *HttpRoutesHandler) HandleGetImages(w http.ResponseWriter, r *http.Request) {
	panic("to be implemented...")
}
