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

// * Params
// * Params is used to store parameter sent from the client
// * Params is NOT query parameter, but instead a parameter in the URL
// * Dyanamic URL changes depending on the parameter, for example /products/1, /products/2
// * /products/:id
// * ServeMux does not support that

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id")
		fmt.Fprint(writer, text)
	})

	// * Example product id 1
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
