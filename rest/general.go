package rest

import (
	"crud.com/packages/rest/implementation"
	"github.com/gorilla/mux"
	"net/http"
)

func DefineEndpoints() {
	router := mux.NewRouter()
	router.HandleFunc("/save/credentials", implementation.SaveCredentials)
	router.HandleFunc("/{username}", implementation.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/{username}", implementation.GetCredential).Methods(http.MethodGet)
	router.HandleFunc("/", implementation.GetAll).Methods(http.MethodGet)
	http.Handle("/", router)
}
