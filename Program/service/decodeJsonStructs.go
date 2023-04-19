package service

var user userStruct = userStruct{}

type userStruct struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var httpHeaderStatus httpHeaderStatusStruct = httpHeaderStatusStruct{}

type httpHeaderStatusStruct struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
