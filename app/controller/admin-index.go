package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model"
)

type AdminIndextPageData struct {
	User model.User
}

func IndexAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("IndexAdminHandler")
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			// GET
			t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-index.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			data := AdminIndextPageData{
				User: currentUser,
			}
			err = t.ExecuteTemplate(w, "layout", data)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
