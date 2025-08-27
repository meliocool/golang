package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// * HTTP is stateless between client and server
// * Stateless means does not store any data or memorize every request from the client
// * This was meant for ease of scalability on the server side
// * But the server also need to memorize a client, for example when a client is logged in, the server must know
// * So that next time the client does not have to login again
// * To do this use Cookie

// * Cookie is a feature in HTTP where server can return a cookie response (key-value)
// * And the client can store the cookie in the web browser
// * For the next request, client will always bring the cookie automatically
// * The server automatically receives the cookie data that the client brings on every request

// * Creating a cookie
// * Cookie is made inside the server and saved/stored in the client web browser
// * use the function http.SetCookie()

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Cookiez"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success Creating a Cookie!")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	// * request.Cookies() -> gets all cookies in a form of []slice
	cookie, err := request.Cookie("X-Cookiez")
	if err != nil {
		fmt.Fprint(writer, "Cookie does not exist!")
	} else {
		fmt.Fprintf(writer, "Cookie Exist! Hello %s", cookie)
	}
}

func TestCookieServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-cookie", GetCookie)
	mux.HandleFunc("/set-cookie", SetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Melio", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	cookie := new(http.Cookie)
	cookie.Name = "X-Cookiez"
	cookie.Value = "Melio"
	request.AddCookie(cookie)
	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
