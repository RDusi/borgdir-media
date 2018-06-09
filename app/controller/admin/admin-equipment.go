package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type AdminEquipmentPageData struct {
	User            benutzer.User
	EquipementListe []equipment.Equipment
}

func EquipmentAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EquipmentAdminHandler")
	fmt.Println("method:", r.Method)
	currentUser := benutzer.User{ID: 0, Benutzername: "Peter Test", BenutzerTyp: "Verleiher"}
	var currentEquipliste []equipment.Equipment
	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-equipment.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	if r.Method == "GET" {
		currentEquipliste, _ = equipment.GetAll()

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

func DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentEquip, _ := equipment.Get(id)
	currentEquip.Delete()
	http.Redirect(w, r, "/admin/equipment", 301)
}
