package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model"
)

type AdminEquipmentPageData struct {
	User            model.User
	EquipementListe []model.Equipment
}

func EquipmentAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	typ := session.Values["type"]
	if typ.(string) != "Verleiher" && typ == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		fmt.Println("EquipmentAdminHandler")
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-equipment.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			currentEquipliste, _ := model.GetAllEquipment()
			data := AdminEquipmentPageData{
				User:            currentUser,
				EquipementListe: currentEquipliste,
			}
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

func DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentEquip, _ := model.GetEquipmentByID(id)
	currentEquip.Delete()
	http.Redirect(w, r, "admin/equipment", http.StatusFound)
}
