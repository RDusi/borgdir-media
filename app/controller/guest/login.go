package guest

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginHandler")
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-login.tmpl", "template/guest/login.tmpl")
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
		// logic part of log in
		benutzername := r.FormValue("benutzername")
		passwort := r.FormValue("passwort")
		fmt.Println("Benutzername : ", benutzername)
		fmt.Println("Passwort: ", passwort)
	}
}
