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
	Mods  [4]Mod
}

type Item struct {
	Name        string
	Description string
}

type Mod struct {
	Name        string
	Description string
	Creator     string
	Image       Image
	Tags        string
}

type Image struct {
	Src string
}

type HomeHandler struct{}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmplDir := "./website/templates"
	files := [...]string{
		tmplDir + "/base.html",
		tmplDir + "/components/slider.html",
		tmplDir + "/home/index.html",
	}

	base := template.New("base.html")

	tmpl := template.Must(base.ParseFiles(files[0], files[1], files[2]))

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

	data.Mods[0] = Mod{
		Name:        "Difficulté ajustable",
		Description: "Permet aux joueurs de modifier les dégâts que vous recevez et infligez aux zombies de manière infinie jusqu’au niveau 15",
		Creator:     "Bodarn",
		Image: Image{
			Src: "/assets/images/7DTD-screenshot-01-img-lg.webp",
		},
		Tags: "Survie et immersion",
	}

	data.Mods[1] = Mod{
		Name:        "Difficulté ajustable",
		Description: "Permet aux joueurs de modifier les dégâts que vous recevez et infligez aux zombies de manière infinie jusqu’au niveau 15",
		Creator:     "Bodarn",
		Image: Image{
			Src: "/assets/images/7DTD-screenshot-01-img-lg.webp",
		},
		Tags: "Survie et immersion",
	}

	data.Mods[2] = Mod{
		Name:        "Difficulté ajustable",
		Description: "Permet aux joueurs de modifier les dégâts que vous recevez et infligez aux zombies de manière infinie jusqu’au niveau 15",
		Creator:     "Bodarn",
		Image: Image{
			Src: "/assets/images/7DTD-screenshot-01-img-lg.webp",
		},
		Tags: "Survie et immersion",
	}

	data.Mods[3] = Mod{
		Name:        "Difficulté ajustable",
		Description: "Permet aux joueurs de modifier les dégâts que vous recevez et infligez aux zombies de manière infinie jusqu’au niveau 15",
		Creator:     "Bodarn",
		Image: Image{
			Src: "/assets/images/7DTD-screenshot-01-img-lg.webp",
		},
		Tags: "Survie et immersion",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
