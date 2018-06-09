package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
)

type AdminIndextPageData struct {
	User benutzer.User
}

func IndexAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IndexAdminHandler")
	fmt.Println("method:", r.Method)
	currentUser := benutzer.User{ID: 0, Benutzername: "Peter Test", BenutzerTyp: "Verleiher"}
	if r.Method == "GET" {
		// GET

		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-index.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		data := AdminIndextPageData{
			User: currentUser,
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
