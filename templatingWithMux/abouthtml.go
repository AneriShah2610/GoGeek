package main

import (
	"net/http"
)

func aboutHTML(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "about.html", nil)
}
