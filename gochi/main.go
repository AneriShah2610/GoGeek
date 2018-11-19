package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// create router
	router := chi.NewRouter()

	router.Use(middleware.RequestID)

	router.Use(middleware.Logger)

	router.Route("/say", func(router chi.Router) {
		router.Get("/", requestSay)
	})
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println("ListenAndServ errror : ", err)
	}

}
func requestSay(writer http.ResponseWriter, reader *http.Request) {
		fmt.Println(writer, "Hello welcome to you!...")
}
