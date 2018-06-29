package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model"
)

type ModClient struct {
	User                 model.User
	EquipmentListeByUser []model.Equipment
}

type AdminClientsPageData struct {
	User      model.User
	UserListe []ModClient
}

func ClientsAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("ClientsAdminHandler")
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			// GET
			t, err := template.ParseFiles("template/layout.tmpl", "template/header-admin-clients.tmpl", "template/admin-clients.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			userListe, err := model.GetAllBenutzer()
			fmt.Println("ASDASDASDASDADDASDAD", userListe)
			var modClient []ModClient
			for _, client := range userListe {
				meineGeraeteListeByUser, _ := model.GetAllMeineGeraeteByUserId(client.ID)
				var equipmentListe []model.Equipment
				for _, meingeraetitem := range meineGeraeteListeByUser {
					equipitem := meingeraetitem.Equipment
					equipmentListe = append(equipmentListe, equipitem)
				}
				var moddedClient ModClient
				moddedClient.User = client
				moddedClient.EquipmentListeByUser = equipmentListe
				modClient = append(modClient, moddedClient)
			}
			data := AdminClientsPageData{
				User:      currentUser,
				UserListe: modClient,
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
