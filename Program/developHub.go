package main

import (
	"BackendTutorial/chapter3/myapp"
	"net/http"
)

const PORT = ":9190"

func main() {

	http.ListenAndServe(PORT, myapp.NewHTTPHandler())

}
