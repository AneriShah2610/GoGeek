package main

import (
	"net/http"
)

func contactHTML(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", nil)
}
