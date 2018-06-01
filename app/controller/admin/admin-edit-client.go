package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
)

type AdminEditClientPageData struct {
	Benutzername string
	BenutzerTyp  string
	UserData     benutzer.User
}

func EditClientAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	currentBenutzerName := "Peter Dieter"
	currentBenutzerTyp := "Benutzer"
	currentUser, _ := benutzer.Get(1)

	if r.Method == "GET" {
		// GET
		data := AdminEditClientPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			UserData:     currentUser,
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

		benutzername := r.FormValue("benutzername")
		email := r.FormValue("email")
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		bild := "../../../static/images/" + handler.Filename

		var passwort string
		if r.FormValue("passwortneu") == r.FormValue("passwortneuwdh") {
			fmt.Println("gleiches Passwort wurde eingegeben")
			passwort = r.FormValue("passwortneu")
		}

		user := benutzer.User{ID: currentUser.ID, Benutzername: benutzername, Email: email, Passwort: passwort, Bild: bild}
		if r.FormValue("speichern") == "1" {
			fmt.Println("Konto sperren")
		}

		if r.FormValue("speichern") == "2" {
			user.Update()
			fmt.Println("User wurde geupdated")
		}
		defer file.Close()
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		data := AdminEditClientPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
		}
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
