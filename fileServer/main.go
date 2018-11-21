package main

import "net/http"

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
}
