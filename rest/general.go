package rest

import (
	"crud.com/packages/rest/implementation"
	"net/http"
)

func DefineEndpoints() {
	http.HandleFunc("/save/credentials", implementation.SaveCredentials)
}
