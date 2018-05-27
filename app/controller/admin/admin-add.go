package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/jhoefker/borgdir-media/app/model/profil"
)

func AddAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddEquipmentAdminHandler")
	fmt.Println("method:", r.Method)

	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-add.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	if r.Method == "GET" {
		// GET
		data := profil.ProfilDummyDataAdmin()
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
		anzahl := r.FormValue("anzahl")
		hinweise := r.FormValue("hinweise")
		speichern := r.FormValue("speichern")

		fmt.Println("Bezeichnung: ", bezeichnung)
		fmt.Println("Kategorie: ", kategorie)
		fmt.Println("Inventar-Nummer: ", invnr)
		fmt.Println("Lagerort: ", lagerort)
		fmt.Println("Inhalt: ", inhalt)
		fmt.Println("Anzahl:", anzahl)
		fmt.Println("Hinweise:", hinweise)
		fmt.Println("Speicher: ", speichern)

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		data := profil.ProfilDummyDataAdmin()
		err = t.ExecuteTemplate(w, "layout", data)
		f, err := os.OpenFile("./static/images/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
