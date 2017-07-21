package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rumyantseva/production-ready-microservices/practice/09-versioning/version"
)

// home returns the path of current request
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("Request received.")
	fmt.Fprintf(w, "Repo: %s, Commit: %s, Version: %s", version.REPO, version.COMMIT, version.RELEASE)
}
