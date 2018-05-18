package guest

import (
	"fmt"
	"html/template"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-register.tmpl", "template/guest/register.tmpl")
		if err != nil {
			fmt.Println(err)
		}
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
		fmt.Println("Benutzername : ", benutzername)
		fmt.Println("E-Mail : ", email)
		fmt.Println("Passwort: ", passwort)
		fmt.Println("Passwort wdh: ", passwortwdh)
	}
}
