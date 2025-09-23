package routes

import (
	// "html/template"
	// "log"

	"log"
	"net/http"
	"text/template"
)

type HomeData struct {
	Items [3]Item
}

type Item struct {
	Name        string
	Description string
}

type HomeHandler struct{}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmplDir := "./website/templates"
	files := [...]string{
		tmplDir + "/base.html",
		tmplDir + "/home/index.html",
	}

	base := template.New("base.html")

	tmpl := template.Must(base.ParseFiles(files[0], files[1]))

	data := HomeData{}
	data.Items[0] = Item{
		Name:        "Difficulté accrue",
		Description: "Nous essayons de monter le niveau de difficulté du jeu afin de se rapprocher le plus possible d’une vraie expérience de survie.",
	}

	data.Items[1] = Item{
		Name:        "Communauté ouverte",
		Description: "Sur Discord, ou en jeu, retrouvez des partenaires de jeu, échangez et partager vos moments forts avec les autres.",
	}

	data.Items[2] = Item{
		Name:        "Equipe à l'écoute",
		Description: "Des salons Discord sont dédiés pour les questions, le support en cas de bugs et nous vous répondrons avec plaisir.",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
