package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/profil"
)

func AddAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddEquipmentAdminHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-add.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		data := profil.ProfilDummyDataAdmin()
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}

	if r.Method == "POST" {
		// POST
		r.ParseForm()
		// logic part of Profil

		bezeichnung := r.FormValue("bezeichnung")
		kategorie := r.FormValue("kategorie")
		invnr := r.FormValue("invnr")
		lagerort := r.FormValue("lagerort")
		inhalt := r.FormValue("inhalt")
		anzahl := r.FormValue("anzahl")
		hinweise := r.FormValue("hinweise")

		fmt.Println("Bezeichnung: ", bezeichnung)
		fmt.Println("Kategorie: ", kategorie)
		fmt.Println("Inventar-Nummer: ", invnr)
		fmt.Println("Lagerort: ", lagerort)
		fmt.Println("Inhalt: ", inhalt)
		fmt.Println("Anzahl:", anzahl)
		fmt.Println("Hinweise:", hinweise)
	}
}
