package server

import (
	"encoding/json"
	"fmt"
	"img-chat-bot/model"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (h *HttpRoutesHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "server is up and running"}`))
}

func (h *HttpRoutesHandler) HandleAddImages(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving the file from form data: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := "./images/"

	// Create the uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		fmt.Println("Error creating directory:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a file with a unique name in the uploads directory
	filePath := filepath.Join(uploadDir, handler.Filename)
	newFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the uploaded file data to the newly created file
	_, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Image uploaded successfully: %s", handler.Filename)
}

func (h *HttpRoutesHandler) HandleGetImages(w http.ResponseWriter, r *http.Request) {
	panic("to be implemented...")
}

func (h *HttpRoutesHandler) HandleUserPrompt(w http.ResponseWriter, r *http.Request) {
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
