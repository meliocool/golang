package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// * ServeMux
// * When creating a web, we would want to make a lot of endpoint URL
// * HandlerFunc cannot support that
// * Alternative implementation of Handler is ServeMux

func TestServeMux(t *testing.T) {
	// * ServeMux is an implementation of handler that supports multiple endpoint
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Welcome!")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hi Web!")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello Web!")
		if err != nil {
			panic(err)
		}
	})

	// * URL Pattern
	// * If you add /images/ then anything after that slash will execute /images/
	// * But if the longer one will be prioritized
	// * For example /images/thumbnails if exist, will be rendered instead of /images/

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Image!")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/images/thumbnails", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Thumbnail!")
		if err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
