package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// * Template Layout
// * Header and Footer are mostly always the same in pages
// * Use layout for repetitive similar layout in each pages
// * Import from other templates
// * {{template "templateName"}} -> without sending any data
// * {{template "templateName" .Value}} -> sends data

func TemplateActionLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml"))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Action Layout",
		"Name":  "Melio",
	})
}

func TestTemplateActionLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// * Template Name
// * If we want to change template name
// * {{define "name"}} TEMPLATE {{end}}
