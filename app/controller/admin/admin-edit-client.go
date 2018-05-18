package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/clients"
)

func EditClientAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		data := clients.CreateClientDummy()
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
		passwortneu := r.FormValue("passwortneu")
		passwortneuwdh := r.FormValue("passwortneuwdh")

		fmt.Println("Benutzername: ", benutzername)
		fmt.Println("E-Mail: ", email)
		fmt.Println("Passwort Neu: ", passwortneu)
		fmt.Println("Passwort Neu Wdh: ", passwortneuwdh)
	}
}
