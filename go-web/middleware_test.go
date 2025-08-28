package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// * Middleware
// * Middleware is a feature to add code before and after a handler is executed
// * There is no Middleware in GoLang
// * We must create our own with The Handler Interface

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Handler Executed")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Handler Executed")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (handler ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error Happened!")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()
	handler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hi Executed")
		fmt.Fprint(writer, "Hello Hi")
	})

	// * Error Handler
	// * If panic happens inside the Handler, we can recover inside the middleware and turn the panic into an error response
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Executed")
		panic("Stop!")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// * ServeMux does not have advanced feature like path variables, auto binding parameter, or proper middleware
// * Therefore many people use routing library
// * Example: httprouter, gorillamux, etc
