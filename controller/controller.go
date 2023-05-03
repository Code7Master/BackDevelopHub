package controller

import ( 
	"BackDevelopHub/service/handler"	
	"net/http"

)
func NewHTTPHandler() http.Handler {
	
	mux := http.NewServeMux()

	// user
	mux.HandleFunc("/auth/singup", handler.AuthSingup)
	mux.HandleFunc("/auth/login", handler.AuthLogin)
	
	// question
	mux.HandleFunc("/question/total", handler.QuestionTotal)
	mux.HandleFunc("/question/preview", handler.QuestionPreview)	// query string
	mux.HandleFunc("/question/detail", handler.QuestionDetail)		// query string
	
	//answer


	return mux
}
