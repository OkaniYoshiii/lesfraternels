package routes

import (
	"html/template"
	"log"
	"net/http"
)

type LoginHandler struct {
}

func (handler *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmplDir := "./website/templates"
	files := []string{
		tmplDir + "/base.html",
		tmplDir + "/login/index.html",
	}

	base := template.New("base.html")
	tmpl := template.Must(base.ParseFiles(files...))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
