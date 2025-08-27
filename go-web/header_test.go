package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// * Header
// * Header is additional information that is sent from client to server or vice versa
// * Header may contain Content-Type, etc
// * To access Header from Request, use Request.Header that returns map[string][]string
// * Header key is NOT case sensitive, unlike query parameter

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	ContentType := request.Header.Get("Content-Type")
	fmt.Fprint(writer, ContentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "MyGoatSamulehu")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("X-Powered-By"))
}
