package implementation

//TODO: handle errors
import (
	"crud.com/packages/configuration"
	"crud.com/packages/domain"
	"crud.com/packages/token"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

const username = "username"

func SaveCredentials(writer http.ResponseWriter, request *http.Request) {
	var crd *domain.CredentialDTO
	reqBody, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Print(err)
		json.NewEncoder(writer).Encode(err)
	}

	json.Unmarshal(reqBody, &crd)
	var cred = &domain.Credential{
		Username: crd.Username,
		Email:    crd.Email,
	}
	domain.SetPassword(cred, crd.UserPassword)
	var value = configuration.Pdatabase.Find(&domain.Credential{Username: crd.Username})
	if value.RowsAffected > 0 {
		print(domain.Update(cred))
		json.NewEncoder(writer).Encode(crd)
	} else {
		print(domain.Save(cred))
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

func GetAll(writer http.ResponseWriter, _ *http.Request) {
	var creds []domain.Credential
	configuration.Pdatabase.Find(&creds)
	response(writer, creds)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credDto *domain.CredentialDTO
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		json.Unmarshal(requestBody, &credDto)
		var cred = domain.Credential{Email: credDto.Email}
		configuration.Pdatabase.Find(&cred)
		err = domain.IsPasswordDiff([]byte(cred.PasswordHash), []byte(credDto.UserPassword))
		if err == nil {
			print("success")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(orchestrateToken(credDto.Email))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("fail")
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
	//Implementation to be done
}

func response(writer http.ResponseWriter, value any) {
	json.NewEncoder(writer).Encode(value)
}

var secretKey = "your-256-bit-secret"

func orchestrateToken(email string) string {
	var duration, err = time.ParseDuration("12h")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	var maker = token.JWTMaker{}
	var jwt, _ = maker.NewJWTMaker(secretKey)

	tkn, err := jwt.CreateToken(email, duration)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return tkn
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello world")
}
