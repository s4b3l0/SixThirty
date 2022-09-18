package implementation

import (
	"crud.com/packages/domain"
	"net/http"
)

func SaveCredentials(writer http.ResponseWriter, request *http.Request) {
	var cred = domain.Credential{
		Username: "",
		Email:    "",
		Password: "",
	} //save stuff
	print(cred.Username + " " + cred.Email + " " + cred.Password)
}
