package main

import (
	"crud.com/packages/configuration"
	"crud.com/packages/rest"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	rest.DefineEndpoints()
	configuration.Connect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
