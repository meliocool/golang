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

// * HTTP Method
// * Like ServeMux, you can also add route to Router
// * The difference is in Router, we can determine the HTTP method that can be used
// * To add it use router.GET(), router.POST(), router.PUT(), etc
// * We use httprouter.Handle(http.ResponseWriter, *http.Request, Params)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello Get")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Get", string(body))
}
