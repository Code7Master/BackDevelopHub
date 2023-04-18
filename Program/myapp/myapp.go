package myapp

import (
	"BackDevelopHub/Program/database"
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

	DB := database.InitDB()
	var data []byte
	if database.NewDatabaseQuery().IsSameUsernameEmail(DB, user.Username, user.Email) {
		// Username과 email 중 하나만이라도 일치했다면
		data, err = json.Marshal(HTTPHeaderStatus{Status: 404, Message: "There are already registered or existing username."})
		if err != nil {
			fmt.Println("[Server] ", err)
		}
		responseWriter.WriteHeader(http.StatusNotFound) // 404

	} else {
		data, err = json.Marshal(HTTPHeaderStatus{Status: 201, Message: "Your membership has been successfully completed."})
		if err != nil {
			fmt.Println("[Server] ", err)

		}
		responseWriter.WriteHeader(http.StatusOK)
		database.NewDatabaseQuery().SingUp(DB, user.Username, user.Email, user.Password)
	}

	responseWriter.Header().Add("content-type", "application/json")
	fmt.Fprint(responseWriter, string(data))

}
