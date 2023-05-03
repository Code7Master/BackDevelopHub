package handler

import (
	"fmt"
	"net/http"
)

func sendJSONResponse(responseWriter http.ResponseWriter, statusCode int, json string) {
	fmt.Println("sendJSONResponse")
	responseWriter.Header().Set("Content-Type", "application/json");
	responseWriter.WriteHeader(statusCode)
	responseWriter.Write([]byte(json))
}