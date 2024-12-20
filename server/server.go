package server

import (
	"encoding/json"
	"fmt"
	"img-chat-bot/constants"
	"img-chat-bot/model"
	"img-chat-bot/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *HttpRoutesHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "server is up and running"}`))
}

func (h *HttpRoutesHandler) HandleAddImages(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	fileData, fileHeader, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving the file from form data: ", err)
		utils.HTTPFailWith4xx(err.Error(), w)
		return
	}
	defer fileData.Close()
	userID := h.getUserIdFromHeaders(r)
	err = h.ChatBot.SaveUserImage(r.Context(), model.FileDetailsModel{
		Header: fileHeader,
		Data:   fileData,
	}, userID)
	if err != nil {
		utils.HTTPFailWith4xx(err.Error(), w)
		return
	}
	utils.HTTPSuccessWith200(fmt.Sprintf("Image uploaded successfully: %s", fileHeader.Filename), w)
}

func (h *HttpRoutesHandler) HandleGetImages(w http.ResponseWriter, r *http.Request) {
	panic("to be implemented...")
}

func (h *HttpRoutesHandler) HandleUserPrompt(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HTTPFailWith4xx(err.Error(), w)
		return
	}

	requestBody := model.RequestModel{}
	err = json.Unmarshal(reqBody, &requestBody)
	if err != nil {
		utils.HTTPFailWith4xx(err.Error(), w)
		return
	}

	ctx := r.Context()
	userID := h.getUserIdFromHeaders(r)
	resp, err := h.ChatBot.GenerateResponse(ctx, requestBody.Prompt, userID)
	if err != nil {
		utils.HTTPFailWith5xx(err.Error(), w)
		return
	}
	utils.HTTPSuccessWith200(resp, w)
	return
}

func (h *HttpRoutesHandler) getUserIdFromHeaders(r *http.Request) int {
	userID := r.Header.Get("user_id")
	if userID == "" {
		return constants.SYSTEM_USER
	}
	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return constants.SYSTEM_USER
	}
	return intUserID
}
