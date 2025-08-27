package go_web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body) // * Use ReadAll to read
	if err != nil {
		return
	}
	bodyString := string(body)
	assert.Equal(t, "Hello World!", bodyString, "Result Must be Hello World!")
	fmt.Println("Testing HTTP Hello World with Assertion Complete!")
}
