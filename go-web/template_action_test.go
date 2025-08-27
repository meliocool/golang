package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// * Template supports command action, like looping, if else, etc

// * if else
// * {{if.Value}} T1 {{end}} If Value is not empty, then T1 is executed, if it is the nothing happens
// * {{if.Value}} T1 {{else}} T2 {{end}} If Value is not empty, then T1 is executed, if it is the T2 is executed
// * {{if.Value1}} T1 {{else if.Value2}} T2 {{else}} T3 {{end}} You get the gist

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data",
		Address: Address{
			Street: "Pancoran",
		},
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// * Comparison Operator
// * eq -> arg1 == arg2
// * ne -> arg1 1= arg2
// * lt -> arg1 < arg2
// * le -> arg1 <= arg2
// * gt -> arg1 > arg2
// * ge -> arg1 >= arg2
// * {{if operator .arg1 .arg2}}
// * THe operator is actually a function

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Action Operator",
		"FinalValue": 90,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// * Range
// * Range is used to iterate through data in an array slice map or channel
// * {{range $index, $element := .Value}} T1 {{end}}
// * {{range $index, $element := .Value}} T1  {{else}} T2 {{end}}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Action Range",
		"Hobbies": []string{
			"Gaming", "Coding", "Drakor",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// * With
// * With changes the scope of dot (.) to an object
// * {{with.Value}} T1 {{end}}
// * {{with.Value}} T1 {{else}} T2 {{end}}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "Action With",
		"Name":  "Melio",
		"Address": map[string]interface{}{
			"Street": "Pancoran",
			"City":   "Jakarta",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
