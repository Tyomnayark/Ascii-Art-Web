package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errorMsg := http.StatusText(status)

	response := struct {
		Error string `json:"error"`
		Code  int    `json:"code"`
	}{
		Error: errorMsg,
		Code:  status,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON marshaling error", http.StatusInternalServerError)
		return
	}

	if _, err2 := w.Write(jsonData); err2 != nil {
		http.Error(w, "JSON writing error", http.StatusInternalServerError)
	}
}
