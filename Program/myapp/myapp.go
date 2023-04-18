package myapp

import (
	"BackDevelopHub/Program/database"
	"BackDevelopHub/Program/log"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

var logger log.Logger = *log.NewLogger()
var databseQuerier database.DatabaseQuerier = *database.NewDatabaseQuerier()
var DB *sql.DB = database.InitDB()

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

	// POST
	router.HandleFunc("/auth/singup", singUp)
	router.HandleFunc("/auth/login", login)
	return router
}

func singUp(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Record("[BackDevelopHub/Program/myapp/myapp.go (func singUp)__GOOD]: start")
	user := new(User)
	err := json.NewDecoder(request.Body).Decode(user)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		logger.Record(fmt.Sprintf("[BackDevelopHub/Program/myapp/myapp.go (func singUp)__ERROR] jsno Decode: %s", err))
		fmt.Fprint(responseWriter, "Bad request: ", err)
		return
	}

	var data []byte
	if databseQuerier.IsSameUsernameOREmail(DB, user.Username, user.Email) {
		// Username과 email 중 하나만이라도 일치했다면
		data, err = json.Marshal(HTTPHeaderStatus{Status: 404, Message: "There are already registered or existing username."})
		if err != nil {
			logger.Record(fmt.Sprintf("[BackDevelopHub/Program/myapp/myapp.go (func singUp)__ERROR] (if IsSameUsernameEmail -> true) json Marshal: %s", err))
		}
		responseWriter.WriteHeader(http.StatusNotFound) // 404

	} else {
		data, err = json.Marshal(HTTPHeaderStatus{Status: 201, Message: "Your membership has been successfully completed."})
		if err != nil {
			logger.Record(fmt.Sprintf("[BackDevelopHub/Program/myapp/myapp.go (func singUp)__ERROR] (if IsSameUsernameEmail -> false) json Marshal: %s", err))
		}
		responseWriter.WriteHeader(http.StatusOK)
		databseQuerier.SingUp(DB, user.Username, user.Email, user.Password)
	}
	responseWriter.Header().Add("content-type", "application/json") // 엔터티 헤더
	fmt.Fprint(responseWriter, string(data))
}

func login(responseWriter http.ResponseWriter, request *http.Request) {
	//TODO:: 로그인 구현
	// user := new(User)
	// err := json.NewDecoder(request.Body).Decode(user)
}
