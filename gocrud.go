package main

import (
	"crud.com/packages/domain"
	"crud.com/packages/rest"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	rest.DefineEndpoints()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func saveCredentials(writer http.ResponseWriter, request *http.Request) {
	var cred = domain.Credential{
		Username: "",
		Email:    "",
		Password: "",
	} //save stuff
	print(cred.Username + " " + cred.Email + " " + cred.Password)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
