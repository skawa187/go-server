package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello page"))
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/hello", helloHandler)

	server := &http.Server{
		Addr:    ":3333",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server error: ", err)
	}
}
