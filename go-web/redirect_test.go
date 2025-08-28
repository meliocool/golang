package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// * Redirect
// * When making a website, usually we use Redirect
// * For example after login, redirect to main page
// * Add response code 3xx and add header Location
// * But GoLang make it easy with http.Redirect

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// * Login
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// * Login
	http.Redirect(writer, request, "https://www.instagram.com/dhitandhitan", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// * Upload File
// * File uploads are often called -> MultiPart
// * In Go, use request.ParseMultipartForm(size) or take the data with request.FormFile(name) -> Automatic Parsing
// * multipart.File -> representation of the file itself
// * multipart.FileHeader -> Information about the file
