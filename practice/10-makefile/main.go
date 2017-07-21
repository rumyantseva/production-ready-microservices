package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

var log = logrus.New()

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/
func main() {
	log.Info("Initialize service...")

	router := httprouter.New()
	router.GET("/", home)

	log.Info("Service is ready to listen and serve.")
	http.ListenAndServe(":8000", router)
}
