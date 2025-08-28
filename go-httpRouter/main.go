package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// * The core of HttpRouter is a struct Router
// * Router is an implementation of http.Handler, which means it can be added easily to http.Server

func main() {
	// * To create a Router, use the function httprouter.New(), that will return Router Pointer
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
