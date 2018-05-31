package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type MyEquipmentPageData struct {
	Benutzername string
	BenutzerTyp  string
	MeineGeraete []equipment.Equipment //noch aendern
	Vorgemerkte  []equipment.Equipment //noch aendern
}

func MyEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MyEquipmentHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-myequip.tmpl", "template/user/my-equipment.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		currentBenutzerName := "Peter Dieter"
		currentBenutzerTyp := "Benutzer"
		meineGeraeteListe, _ := equipment.GetAll() // noch aendern zu DB
		data := MyEquipmentPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			MeineGeraete: meineGeraeteListe,
			Vorgemerkte:  []equipment.Equipment{},
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
