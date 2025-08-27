package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// * Form POST
// * Submitting forms uses either GET or POST
// * The form data will be sent in Query Parameter if the method is GET
// * The form data will be sent via HTTP request body if the method is POST
// * In GoLang, just use Request.PostForm
// * But before taking the data from PostForm, call the method Request.ParseForm()

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	// * request.PostFormValue("first_name") automatic parsing

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")
	fmt.Fprintf(writer, "Hi %s%s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Melio&last_name=das")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	recorder := httptest.NewRecorder()
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
