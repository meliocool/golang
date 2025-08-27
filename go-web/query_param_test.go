package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// * Query parameter is a feature that can be used when creating a Web
// * Query parameter is used to send data from client to server
// * Query parameter is placed inside the URL
// * To add query parameter use ?name=value inside the URL
// * URL representation in GoLang is url.URL
// * In the request parameter there is an attribute URL that is filled with data url.URL
// * From the URL Data, we can take the query parameter data that is sent from the client with method Query() that will return a Map

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Melio", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// * Multiple Query Parameter
// * with &

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s%s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Melio&last_name=das", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// * Multiple Value Parameter
// * URL does a parsing for query parameter and stores it in map[string][]string
// * That means in 1 key query parameter, there can be multiple value
// * Use name=Melio&name=das

func MultipleValueParameter(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprintln(writer, strings.Join(names, " "))
}

func TestMultipleValueParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Melio&name=das", nil)
	recorder := httptest.NewRecorder()

	MultipleValueParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
