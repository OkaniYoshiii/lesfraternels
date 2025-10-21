package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type LoginHandler struct {
}

func (handler *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		validate := validator.New(validator.WithRequiredStructEnabled())

		password := r.PostFormValue("login_password")
		email := r.PostFormValue("login_email")

		if err := validate.Var(password, "required"); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		if err := validate.Var(email, "required,email"); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}

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
