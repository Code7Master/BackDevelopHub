package main

import (
	"BackDevelopHub/Program/myapp"
	"net/http"
)

const PORT = ":9190"

func main() {
	// 아래의 코드가 먼저 실행되게 하기
	// 	log.NewLogger().Record("[DevelopHub] Server Start!!")
	http.ListenAndServe(PORT, myapp.NewHTTPHandler())
}
