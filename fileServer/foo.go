package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("userfile")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error while uplodaing file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error while Reading file", http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `<form action="/" method="post" enctype="multipart/form-data">
    upload a file<br>
    <input type="file" name="userfile"><br>
    <input type="submit">
	</form>
	<br>
	<br>
	<h1>%v</h1>`, s)
}
