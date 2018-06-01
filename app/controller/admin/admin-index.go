package admin

import (
	"fmt"
	"html/template"
	"net/http"
)

type AdminIndextPageData struct {
	Benutzername string
	BenutzerTyp  string
}

func IndexAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IndexAdminHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET

		currentBenutzerName := "Peter Meier"
		currentBenutzerTyp := "Verleiher"

		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-index.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		data := AdminIndextPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
