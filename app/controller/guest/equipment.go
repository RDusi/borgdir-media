package guest

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type EquipmentPageData struct {
	Benutzername   string
	BenutzerTyp    string
	EquipmentListe []equipment.Equipment
}

func EquipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EquipmentHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-equip.tmpl", "template/guest/equipment.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		currentBenutzerName := "Peter Dieter"
		currentBenutzerTyp := "Benutzer"
		equipmentListe, err := equipment.GetAll()
		data := EquipmentPageData{
			Benutzername:   currentBenutzerName,
			BenutzerTyp:    currentBenutzerTyp,
			EquipmentListe: equipmentListe,
		}
		fmt.Println("Equipment: ", data)

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
