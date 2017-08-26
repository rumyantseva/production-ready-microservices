package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

var log = logrus.New()

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/
func main() {
	log.Info("Initialize service...")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Required parameter service port is not set")
	}

	router := httprouter.New()
	router.GET("/", home)

	log.Info("Service is ready to listen and serve.")
	http.ListenAndServe(":"+port, router)
}
