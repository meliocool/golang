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

// * Middleware
// * Still gotta create your own Middleware

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Request")
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello Middleware")
	})

	middleware := LogMiddleware{Handler: router}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Middleware", string(body))
}
