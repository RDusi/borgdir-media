package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
)

type AdminClientsPageData struct {
	Benutzername string
	BenutzerTyp  string
	UserListe    []benutzer.User
}

func ClientsAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ClientsAdminHandler")
	fmt.Println("method:", r.Method)

	currentBenutzerName := "Peter Dieter"
	currentBenutzerTyp := "Benutzer"

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-clients.tmpl", "template/admin/admin-clients.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		userListe, err := benutzer.GetAll()
		data := AdminClientsPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			UserListe:    userListe,
		}
		fmt.Println("User: ", data)
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
