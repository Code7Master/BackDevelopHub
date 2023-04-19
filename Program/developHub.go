package main

import (
	"BackDevelopHub/Program/controller"
	"net/http"
)

const PORT = ":9190"

func main() {
	http.ListenAndServe(PORT, controller.NewHTTPHandler())
}
