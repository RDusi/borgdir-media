package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/categorie"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
	"github.com/jhoefker/borgdir-media/app/model/storage"
)

type AdminAddPageData struct {
	User benutzer.User
}

func init() {
	tmpl, err = template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-add.tmpl")
	if err != nil {
		fmt.Println(err)
	}
}

func AddAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddEquipmentAdminHandler")
	fmt.Println("method:", r.Method)
	currentUser := benutzer.User{ID: 0, Benutzername: "Peter Test", BenutzerTyp: "Verleiher"}

	if r.Method == "GET" {
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

		categorie, _ := categorie.Get(kategorieid)
		lagerort, _ := storage.Get(lagerortid)
		equipment := equipment.Equipment{Bezeichnung: bezeichnung, Kategorie: categorie, InventarNr: invnr, Lagerort: lagerort, Inhalt: inhalt, Anzahl: anzahl, Hinweise: hinweise, Bild: bild}
		fmt.Println(equipment)

		defer file.Close()
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		if r.FormValue("speichern") == "2" {
			equipment.Add()
			http.Redirect(w, r, "/admin/add", 301)
		}
	}
}
