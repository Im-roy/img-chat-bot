package utils

import (
	"encoding/json"
	"net/http"
	"path/filepath"
)

func ExtractExtension(fileName string) string {
	extension := filepath.Ext(fileName)
	if extension != "" {
		extension = extension[1:]
	}
	return extension
}

// Response represents a generic response structure
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func WriteResponse(response Response, httpStatusCode int, rw http.ResponseWriter) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatusCode)
	rw.Write(responseJson)
}

func HTTPSuccessWith200(data interface{}, rw http.ResponseWriter) {
	response := Response{
		Success: true,
		Data:    data,
	}
	WriteResponse(response, http.StatusOK, rw)
}

func HTTPFailWithStatusCode(errorMessage string, httpStatusCode int, rw http.ResponseWriter) {
	response := Response{
		Success:      false,
		ErrorMessage: errorMessage,
	}
	WriteResponse(response, httpStatusCode, rw)
}

func HTTPFailWith4xx(errorMessage string, rw http.ResponseWriter) {
	HTTPFailWithStatusCode(errorMessage, http.StatusBadRequest, rw)
}

func HTTPFailWith5xx(errorMessage string, rw http.ResponseWriter) {
	HTTPFailWithStatusCode(errorMessage, http.StatusInternalServerError, rw)
}
