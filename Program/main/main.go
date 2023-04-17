package main

import (
	"BackDevelopHub/Program/myapp"
	"net/http"
)

const PORT = ":9190"

func main() {
	http.ListenAndServe(PORT, myapp.NewHTTPHandler())
}
