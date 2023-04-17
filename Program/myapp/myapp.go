package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HTTPHeaderStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewHTTPHandler() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/auth/singup", singUp)

	return router
}

func singUp(responseWriter http.ResponseWriter, request *http.Request) {
	user := new(User)
	err := json.NewDecoder(request.Body).Decode(user)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(responseWriter, "Bad request: ", err)
		return
	}

	// TODO : DB 값 있는지 없는지 확인하기.

	var data []byte
	if true {
		data, err = json.Marshal(HTTPHeaderStatus{Status: 201, Message: "Your membership has been successfully completed."})
		if err != nil {
			fmt.Println("[Server] ", err)

		}
		responseWriter.WriteHeader(http.StatusOK)

	} else {
		data, err = json.Marshal(HTTPHeaderStatus{Status: 400, Message: "There are already registered or existing username."})
		if err != nil {
			fmt.Println("[Server] ", err)

		}
		// TODO: responseWriter.WriteHeader(http.) 요청이 없다는 http 코드 반환
	}

	responseWriter.Header().Add("content-type", "application/json")
	fmt.Fprint(responseWriter, string(data))

}
