package main

import (
	"log"
	"net/http"

	"github.com/andrearcaina/echo/internal/api/routes"
)

func main() {
	handler := routes.InitAPI()

	server := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}

	log.Printf("Starting server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
