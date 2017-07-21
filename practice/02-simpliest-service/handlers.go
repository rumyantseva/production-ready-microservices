package main

import (
	"fmt"
	"net/http"
)

// home returns the path of current request
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Processing URL %s...\n", r.URL.Path)
}
