package implementation

import (
	"crud.com/packages/domain"
	"fmt"
	"net/http"
)

func SaveCredentials(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf(domain.Something)
	var cred = domain.Credential{
		Username: "",
		Email:    "",
		Password: "",
	} //save stuff
	print(cred.Username + " " + cred.Email + " " + cred.Password)
}
