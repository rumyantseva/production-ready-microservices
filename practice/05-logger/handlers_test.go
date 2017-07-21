package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// TestHandler is the simplest test: check home (/) URL
// go test
func TestHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", home)

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedGreeting := "Processing URL /..."
	testingGreeting := strings.Trim(string(greeting), " \n")
	if testingGreeting != expectedGreeting {
		t.Fatalf(
			"Wrong greeting '%s', expected '%s'",
			testingGreeting, expectedGreeting,
		)
	}
}
