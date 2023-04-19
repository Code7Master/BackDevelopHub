package service

import (
	"encoding/json"
	"net/http"
)

// [/auth/singup]
func SingUp(responseWriter http.ResponseWriter, request *http.Request) {
	err := json.NewDecoder(request.Body).Decode(user)

	err != nil {
		// TODO: Recoder send!
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(responseWriter, "Bad request: ", err)
	}

	// TODO: 밑에 계속 구현
}

// [auth/login]
func Login(responseWriter http.ResponseWriter, request *http.Request) {

}
