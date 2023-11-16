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
	router.HandleFunc("/login", implementation.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", implementation.Logout).Methods(http.MethodPost)
	router.HandleFunc("/all", implementation.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/", implementation.HelloWorld).Methods(http.MethodGet)
	router.HandleFunc("/message", implementation.GetMessage).Methods(http.MethodGet)
	router.HandleFunc("/message", implementation.CreateMessage).Methods(http.MethodPost)
	http.Handle("/", router)
}
