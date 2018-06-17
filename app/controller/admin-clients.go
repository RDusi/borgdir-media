package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model"
)

type AdminClientsPageData struct {
	User      model.User
	UserListe []model.User
}

func ClientsAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	typ := session.Values["type"]
	if typ.(string) != "Verleiher" && typ == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		fmt.Println("ClientsAdminHandler")
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			// GET
			t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-clients.tmpl", "template/admin/admin-clients.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			userListe, err := model.GetAllBenutzer()
			data := AdminClientsPageData{
				User:      currentUser,
				UserListe: userListe,
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
}
