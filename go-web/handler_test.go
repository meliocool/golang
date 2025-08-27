package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// * Handler
// * To receive HTTP Request that is entering the Web Server, there must be a Handler
// * Handler is an interface, where in its contract there must be a function called ServeHTTP()
// * HandlerFunc is an implementation of the Handler interface

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// * Web Logic
		_, err := fmt.Fprint(writer, "Hi Web!")
		if err != nil {
			panic(err)

		}
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// * Request
// * Request is a struct that represents the HTTP Request sent by the Web Browser
// * All request information that is sent, can be obtained in Request
// * For example URL, http method, http header, http body, etc

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
