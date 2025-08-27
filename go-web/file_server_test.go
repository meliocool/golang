package go_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// * File Server
// * File Server allows making handler that can be used as static File Server
// * Avoid manually loading File
// * Is a handler, therefore can be added into Server or Mux

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// * File server reads URL, then read the file based on URL
	// * if /static/ -> /resources/static/index.js
	// * use http.StripPrefix()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// * Using Go Embed
// * To embed static files into the exe

//go:embed resources
var resources embed.FS

func TestFileServerGoLangEmbed(t *testing.T) {
	// * folder resources is included in the prefix, so it sucks
	// * Use fs.sub()

	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
