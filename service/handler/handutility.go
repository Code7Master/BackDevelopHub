package handler

import (
	"fmt"
	"net/http"
)

func sendJSONResponse(responseWriter http.ResponseWriter, statusCode int, json []byte) {
	responseWriter.Header().Set("Content-Type", "application/json");
	responseWriter.WriteHeader(statusCode)
	responseWriter.Write(json)
}

func atoi(str string) int {
	var resVal int
	strLen := len(str)

	for i := 0; i < strLen; i++ {
		resVal += int(str[i]) - '0'
		if i+1 != strLen {
			resVal *=  10
		}
	}
	return resVal
}