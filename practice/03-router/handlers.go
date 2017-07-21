package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// home returns the path of current request
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Processing URL %s...\n", r.URL.Path)
}
