package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

func LandingPageHandler(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("templates/landing_page.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "My Go Webpage",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
	}
}