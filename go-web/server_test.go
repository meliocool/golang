package go_web

import (
	"net/http"
	"testing"
)

// * Server is a struct available in package net/http that is used as a representation of Web Server in Go
// * To create a web we must create a Server

func TestServer(t *testing.T) {
	// * When creating Server, we must decide the host and the port on where the Web is running
	server := http.Server{
		Addr: "localhost:8080",
	}

	// * After creating a server, we can run the Server with the function ListenAndServe()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
