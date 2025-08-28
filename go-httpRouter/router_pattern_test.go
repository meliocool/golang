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

// * Router Patterns

// * Named Parameter
// * /user/:user
// * Pattern /user/nate /user/you -> both matches
// * Pattern /user/nate/profile /user/ -> both does not match
func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

// * Catch All Parameter
// * Catches ALL parameter in the url with * symbol
// * Catch All Parameter must be at the end of the URL
// * /src/*filepath
// * Pattern /src/ -> No match
// * Pattern /src/somefile /src/subdir/somefile -> both matches
func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/images/small/gyjicon.jpg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/gyjicon.jpg", string(body))
}
