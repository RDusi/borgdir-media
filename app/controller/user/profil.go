package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/profil"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-profil.tmpl", "template/user/profil.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		data := profil.ProfilDummyDataUser()
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}

	if r.Method == "POST" {
		// POST
		r.ParseForm()
		// logic part of Profil

		benutzername := r.FormValue("benutzername")
		email := r.FormValue("email")
		passwortalt := r.FormValue("passwortalt")
		passwortneu := r.FormValue("passwortneu")
		passwortneuwdh := r.FormValue("passwortneuwdh")
		loeschen := r.FormValue("loeschen")

		fmt.Println("Benutzername: ", benutzername)
		fmt.Println("E-Mail: ", email)
		fmt.Println("Passwort Alt: ", passwortalt)
		fmt.Println("Passwort Neu: ", passwortneu)
		fmt.Println("Passwort Neu Wdh: ", passwortneuwdh)
		fmt.Println("loeschen :", loeschen)
	}
}
