package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// this function is doing render html files in template directory
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsTemplat, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsTemplat.Execute(w, nil)
	if err != nil {
		fmt.Println("Error has been reached", err)

	}
}
