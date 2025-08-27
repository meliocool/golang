package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// * Response Code
// * There are a ton of HTTP Response status Code
// * 1. Informational response (100 - 199)
// * 2. Successful response (200 - 299)
// * 3. Redirects (300 - 399)
// * 4. Client errors (400 - 499)
// * 5. Server errors (500 - 599)
// * By default the response code is 200 or OK
// * Use function ResponseWriter.WriteHeader(int)
// * THere are also variables that represents the status code

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/?name=Melio", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
