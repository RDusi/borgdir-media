package controller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model"
)

type AdminEditEquipmentPageData struct {
	User          model.User
	Equipment     model.Equipment
	AllCategories []model.Categorie
	AllLagerorte  []model.Storage
}

func AdminEditEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("EditClientHandler")
		fmt.Println("method:", r.Method)

		var currentUserEdit model.User
		if r.Method == "GET" {

			tmpl, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-equipment.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			id, _ := strconv.Atoi(r.FormValue("id"))
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]

			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			currentEquipmentEdit, _ := model.GetEquipmentByID(id)
			alleKategorien, _ := model.GetAllKategorien()
			alleLagerorte, _ := model.GetAllStorage()
			fmt.Println("CURRENTUSER: ", currentUserEdit)

			data := AdminEditEquipmentPageData{
				User:          currentUser,
				Equipment:     currentEquipmentEdit,
				AllCategories: alleKategorien,
				AllLagerorte:  alleLagerorte,
			}
			err = tmpl.ExecuteTemplate(w, "layout", data)
			if err != nil {
				fmt.Println(err)
			}
		}

		if r.Method == "POST" {

			r.ParseForm()
			r.ParseMultipartForm(32 << 20)
			// logic part of Profil
			bezeichnung := r.FormValue("bezeichnung")
			kategorieid, _ := strconv.Atoi(r.FormValue("kategorie"))
			invnr := r.FormValue("invnr")
			lagerortid, _ := strconv.Atoi(r.FormValue("lagerort"))
			inhalt := r.FormValue("inhalt")
			anzahl, _ := strconv.Atoi(r.FormValue("anzahl"))
			hinweise := r.FormValue("hinweise")
			file, handler, err := r.FormFile("uploadfile")
			speichern, _ := strconv.Atoi(r.FormValue("speichern"))
			bild := "../../../static/images/" + handler.Filename

			equipmentEdit, _ := model.GetEquipmentByID(speichern)
			if handler.Filename == "" {
				bild = equipmentEdit.Bild
			}

			if r.FormValue("speichern") == strconv.Itoa(equipmentEdit.ID) {
				equipmentEdit.Bezeichnung = bezeichnung
				equipmentEdit.Kategorie, _ = model.GetKategorieById(kategorieid)
				equipmentEdit.InventarNr = invnr
				equipmentEdit.Lagerort, _ = model.GetStorageByID(lagerortid)
				equipmentEdit.Inhalt = inhalt
				equipmentEdit.Anzahl = anzahl
				equipmentEdit.Hinweise = hinweise
				equipmentEdit.Bild = bild
				equipmentEdit.Update()
				fmt.Println("Equipment wurde geupdatet")
				http.Redirect(w, r, "/admin/edit-equipment?id="+strconv.Itoa(equipmentEdit.ID)+"", http.StatusFound)
			}

			defer file.Close()
			fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
			f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		}
	}
}
