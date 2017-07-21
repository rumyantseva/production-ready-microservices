package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/k8s-community/utils/shutdown"
)

var log = logrus.New()

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/
func main() {
	log.Info("Initialize service...")

	router := httprouter.New()
	router.GET("/", home)

	router.GET("/healthz", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Ok")
	})
	router.GET("/readyz", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Ok")
	})

	log.Info("Register graceful shutdown handler...")
	logger := log.WithField("event", "shutdown")
	sdHandler := shutdown.NewHandler(logger)
	go sdHandler.RegisterShutdown(sd)

	log.Info("Service is ready to listen and serve.")
	http.ListenAndServe(":8000", router)
}

// sd does graceful shutdown of the service
func sd() (string, error) {
	// if service has to finish some tasks before shutting down, these tasks must be finished her
	return "Ok", nil
}
