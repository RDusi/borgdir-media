package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type AdminEquipmentPageData struct {
	Benutzername    string
	BenutzerTyp     string
	EquipementListe []equipment.Equipment
}

func EquipmentAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EquipmentAdminHandler")
	fmt.Println("method:", r.Method)

	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-equipment.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	currentBenutzerName := "Peter Dieter"
	currentBenutzerTyp := "Benutzer"
	currentEquipliste, _ := equipment.GetAll()

	if r.Method == "GET" {
		// GET

		data := AdminEquipmentPageData{
			Benutzername:    currentBenutzerName,
			BenutzerTyp:     currentBenutzerTyp,
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
		loeschen := r.FormValue("loeschen")
		bearbeiten := r.FormValue("edit")

		fmt.Println("Loeschen: ", loeschen)
		fmt.Println("Bearbeiten: ", bearbeiten)

		data := AdminEquipmentPageData{
			Benutzername:    currentBenutzerName,
			BenutzerTyp:     currentBenutzerTyp,
			EquipementListe: currentEquipliste,
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-equipment.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	currentEquip, _ := equipment.Get(id)
	currentEquip.Delete()

	currentBenutzerName := "Peter Dieter"
	currentBenutzerTyp := "Benutzer"
	currentEquipliste, _ := equipment.GetAll()

	data := AdminEquipmentPageData{
		Benutzername:    currentBenutzerName,
		BenutzerTyp:     currentBenutzerTyp,
		EquipementListe: currentEquipliste,
	}
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Println(err)
	}
}
