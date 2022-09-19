package implementation

import (
	"crud.com/packages/configuration"
	"crud.com/packages/domain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

const username = "username"

func SaveCredentials(writer http.ResponseWriter, request *http.Request) {
	var crd *domain.Credential
	reqBody, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Print(err)
		json.NewEncoder(writer).Encode(err)
	}

	json.Unmarshal(reqBody, &crd)
	var value = configuration.Pdatabase.Find(&domain.Credential{Username: crd.Username})
	if value.RowsAffected > 0 {
		configuration.Pdatabase.Updates(*crd)
		json.NewEncoder(writer).Encode(crd)
	} else {
		configuration.Pdatabase.Create(crd)
		json.NewEncoder(writer).Encode(crd)
	}
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)
	var username = params[username]
	if len(username) > 0 {
		configuration.Pdatabase.Delete(&domain.Credential{
			Username: username,
		})
	}
}

func GetCredential(writer http.ResponseWriter, request *http.Request) {
	var params = mux.Vars(request)
	var username = params[username]

	if len(username) > 0 {
		var cred, _ = configuration.Pdatabase.Get(username)
		response(writer, cred)
		return
	}
}

func GetAll(writer http.ResponseWriter, r *http.Request) {
	var creds []domain.Credential
	configuration.Pdatabase.Find(&creds)
	response(writer, creds)
}

func response(writer http.ResponseWriter, value any) {
	json.NewEncoder(writer).Encode(value)
}
