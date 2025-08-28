package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// * Method Not Allowed Handler
// * ServeMux cannot do this, because all http method is accepted there
// * But with Router, we can determine the HTTP Method first, if the client sends the wrong method
// * Method Not Allowed error will happen

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Not Allowed!")
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "POST")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Not Allowed!", string(body))
}
