package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type AdminAddPageData struct {
	Benutzername string
	BenutzerTyp  string
}

func AddAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddEquipmentAdminHandler")
	fmt.Println("method:", r.Method)

	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-add.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	if r.Method == "GET" {
		// GET
		currentBenutzerName := "Peter Meier"
		currentBenutzerTyp := "Verleiher"
		data := AdminAddPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}

	if r.Method == "POST" {
		// POST
		r.ParseForm()
		r.ParseMultipartForm(32 << 20)
		// logic part of Profil

		bezeichnung := r.FormValue("bezeichnung")
		kategorie := r.FormValue("kategorie")
		invnr := r.FormValue("invnr")
		lagerort := r.FormValue("lagerort")
		inhalt := r.FormValue("inhalt")
		anzahl, _ := strconv.Atoi(r.FormValue("anzahl"))
		hinweise := r.FormValue("hinweise")

		fmt.Println("Bezeichnung: ", bezeichnung)
		fmt.Println("Kategorie: ", kategorie)
		fmt.Println("Inventar-Nummer: ", invnr)
		fmt.Println("Lagerort: ", lagerort)
		fmt.Println("Inhalt: ", inhalt)
		fmt.Println("Anzahl:", anzahl)
		fmt.Println("Hinweise:", hinweise)

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		if r.FormValue("speichern") == "2" {
			equipment := equipment.Equipment{Bezeichnung: bezeichnung, Kategorie: kategorie, InventarNr: invnr, Lagerort: lagerort, Inhalt: inhalt, Anzahl: anzahl, Hinweise: hinweise, Bild: "../../../static/images/" + handler.Filename + ".jpg"}
			fmt.Println(equipment)
			fmt.Println("Neues Equipment wurde hinzugefuegt")
		}
		currentBenutzerName := "Peter Meier"
		currentBenutzerTyp := "Verleiher"
		data := AdminAddPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
		}
		err = t.ExecuteTemplate(w, "layout", data)
		f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
