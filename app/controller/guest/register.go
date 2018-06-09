package guest

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
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
		user := benutzer.User{Benutzername: benutzername, Email: email, Passwort: passwort, Bild: bild}
		if passwort == passwortwdh {
			fmt.Println("gleiches PW")
			user.Add()
		} else {
			fmt.Println("Passwort stimmt nicht überein")
		}
		err = t.ExecuteTemplate(w, "layout", "data")
		if err != nil {
			fmt.Println(err)
		}

	}
}
