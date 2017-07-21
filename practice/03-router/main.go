package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/
func main() {
	router := httprouter.New()
	router.GET("/", home)

	http.ListenAndServe(":8000", router)
}
