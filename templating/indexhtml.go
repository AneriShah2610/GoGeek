package main

import (
	"net/http"
)

func indexHTML(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "main.html", nil)
}
