package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterHandler")
	fmt.Println("method:", r.Method)

	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-register.tmpl", "template/guest/register.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	if r.Method == "GET" {
		// GET
		err = t.ExecuteTemplate(w, "layout", "data")
		if err != nil {
			fmt.Println(err)
		}
	}
	if r.Method == "POST" {
		// POST
		r.ParseForm()
		// logic part of Register
		benutzername := r.FormValue("benutzername")
		email := r.FormValue("email")
		passwort := r.FormValue("passwort")
		passwortwdh := r.FormValue("passwortwdh")
		bild := "http://via.placeholder.com/350x350"
		user := model.User{Benutzername: benutzername, Email: email, Passwort: passwort, Bild: bild, BenutzerTyp: "Benutzer", AktivBis: "erstmal soweit"}
		if passwort == passwortwdh {
			fmt.Println("gleiches PW")
			user.Add()
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			t, err = template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-register.tmpl", "template/guest/register-falschesPW.tmpl")
			if err != nil {
				fmt.Println(err)
			}
		}
		err = t.ExecuteTemplate(w, "layout", "data")
		if err != nil {
			fmt.Println(err)
		}

	}
}
