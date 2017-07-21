package main

import (
	"net/http"
)

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/any
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}
