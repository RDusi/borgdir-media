package controller

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model"
)

type AdminAddPageData struct {
	User model.User
}

func AddAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("AddEquipmentAdminHandler")
		fmt.Println("method:", r.Method)
		session, _ := store.Get(r, "session")
		benutzername := session.Values["username"]
		fmt.Println(benutzername)
		currentUser, _ := model.GetUserByUsername(benutzername.(string))

		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("template/layout.tmpl", "template/header-admin-std.tmpl", "template/admin-add.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			data := AdminAddPageData{
				User: currentUser,
			}
			err = tmpl.ExecuteTemplate(w, "layout", data)
			if err != nil {
				fmt.Println(err)
			}
		}

		if r.Method == "POST" {
			r.ParseForm()
			r.ParseMultipartForm(32 << 20)

			bezeichnung := r.FormValue("bezeichnung")
			kategorieid, _ := strconv.Atoi(r.FormValue("kategorie"))
			invnr := r.FormValue("invnr")
			lagerortid, _ := strconv.Atoi(r.FormValue("lagerort"))
			inhalt := r.FormValue("inhalt")
			anzahl, _ := strconv.Atoi(r.FormValue("anzahl"))
			hinweise := r.FormValue("hinweise")
			file, handler, _ := r.FormFile("uploadfile")
			bild := "../../../static/images/" + handler.Filename

			categorie, _ := model.GetKategorieById(kategorieid)
			lagerort, _ := model.GetStorageByID(lagerortid)
			equipment := model.Equipment{Bezeichnung: bezeichnung, Kategorie: categorie, InventarNr: invnr, Lagerort: lagerort, Inhalt: inhalt, Anzahl: anzahl, Hinweise: hinweise, Bild: bild}
			fmt.Println(equipment)
			var temp int = 1
			if r.FormValue("speichern") == "2" {
				equipment.Add()
				fmt.Println("Equipment wurde hinzugefügt")
				temp = 2
			}
			defer file.Close()
			fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
			f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			var wert int64 = 0
			wert, _ = io.Copy(f, file)
			if wert != 0 && temp == 2 {
				log.Println("hier")
				http.Redirect(w, r, "/admin/add", http.StatusFound)
			}
		}
	}

}
