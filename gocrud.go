package main

import (
	"crud.com/packages/configuration"
	"crud.com/packages/rest"
	"log"
	"net/http"
)

var connectDb = false

func main() {
	rest.DefineEndpoints()
	if connectDb {
		configuration.Connect()
	}

	port := "5001"
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
