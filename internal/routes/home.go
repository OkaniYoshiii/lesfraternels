package routes

import (
	// "html/template"
	// "log"

	"log"
	"net/http"
	"text/template"
)

type HomeHandler struct{}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmplDir := "./website/templates"
	files := [...]string{
		tmplDir + "/base.html",
		tmplDir + "/home/index.html",
	}

	base := template.New("base.html")

	tmpl := template.Must(base.ParseFiles(files[0], files[1]))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
