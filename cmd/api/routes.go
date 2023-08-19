package main

import (
	"ascii-art-web-dockerize/internal"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

func (app *application) mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.errorResponse(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		tmp, err := template.ParseFiles("templates/404.html")
		if err != nil {
			app.errorResponse(w, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		tmp.Execute(w, "")
	} else {
		templ, err := template.ParseFiles("templates/index.html")
		if err != nil {
			app.errorResponse(w, http.StatusInternalServerError)
			return
		}

		if err := templ.Execute(w, ""); err != nil {
			app.errorResponse(w, http.StatusInternalServerError)
			return
		}
	}
}

func (app *application) artHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.errorResponse(w, http.StatusMethodNotAllowed)
		return
	}
	if r.ContentLength == 0 {
		app.errorResponse(w, http.StatusBadRequest)
		return
	}

	var input struct {
		Text string `json:"text"`
	}
	err1 := json.NewDecoder(r.Body).Decode(&input)
	if err1 != nil {
		app.errorResponse(w, http.StatusBadRequest)
		return
	}

	selectedStyle := r.URL.Query().Get("style")
	if len(selectedStyle) == 0 {
		selectedStyle = "Standard"
	}
	asciiArt, err2 := internal.GenerateASCII(input.Text, selectedStyle)

	if err2 != nil {
		app.errorResponse(w, http.StatusInternalServerError)
		return
	}

	response := struct {
		ASCIIArt string `json:"asciiArt"`
	}{
		ASCIIArt: asciiArt,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err4 := json.NewEncoder(w).Encode(response)
	if err4 != nil {
		app.errorResponse(w, http.StatusInternalServerError)
		return
	}

}
func (app *application) saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.errorResponse(w, http.StatusMethodNotAllowed)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		app.errorResponse(w, http.StatusBadRequest)
		return
	}
	selectedStyle := r.Form.Get("style")
	if len(selectedStyle) == 0 {
		selectedStyle = "Standard"
	}
	asciiArt, err2 := internal.GenerateASCII(content, selectedStyle)

	if err2 != nil {
		app.errorResponse(w, http.StatusInternalServerError)
		return
	}
	filePath := "./ascii_art.txt"

	err := ioutil.WriteFile(filePath, []byte(asciiArt), 0644)
	if err != nil {
		app.errorResponse(w, http.StatusInternalServerError)
		return
	}

	fileName := "ascii_art.txt"
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	w.Header().Set("Content-Type", "text/plain;charset=utf-8")

	w.WriteHeader(http.StatusOK)
}
