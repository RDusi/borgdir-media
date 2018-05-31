package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type CartPageData struct {
	Benutzername   string
	BenutzerTyp    string
	Ausleihvorgang []equipment.Equipment //noch aendern
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CartHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-cart.tmpl", "template/user/cart.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		currentBenutzerName := "Peter Dieter"
		currentBenutzerTyp := "Benutzer"
		data := CartPageData{
			Benutzername:   currentBenutzerName,
			BenutzerTyp:    currentBenutzerTyp,
			Ausleihvorgang: []equipment.Equipment{},
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}

	if r.Method == "POST" {
		// POST
		r.ParseForm()
		// logic part of Equipment
	}
}
