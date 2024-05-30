package main

import (
	"log"
	"net/http"

	"github.com/andrearcaina/echo/internal/api/routes"
)

func main() {
	// initializes a http multiplexer (basically a request router)
	router := routes.InitRouter()

	// initializes a http server with the address and given multiplexer
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	// logs the server address and starts the server
	log.Printf("Starting server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
