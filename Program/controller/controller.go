package controller

import (
	"BackDevelopHub/Program/service"
	"net/http"
)

func NewHTTPHandler() http.Handler {

	router := http.NewServeMux()
	router.HandleFunc("/auth/singup", service.SingUp)
	router.HandleFunc("/auth/login", service.Login)
	return router
}
