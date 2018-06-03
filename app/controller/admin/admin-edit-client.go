package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
)

type AdminEditClientPageData struct {
	Benutzername string
	BenutzerTyp  string
	UserData     benutzer.User
}

var tmpl *template.Template
var err error

// Is executed automatically on package load
func init() {
	tmpl, err = template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client.tmpl")
	if err != nil {
		fmt.Println(err)
	}
}

func EditClientAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	id, _ := strconv.Atoi(r.FormValue("id"))
	currentBenutzerName := "Peter Dieter"
	currentBenutzerTyp := "Benutzer"
	currentUser, _ := benutzer.Get(id)

	if r.Method == "GET" {
		// GET
		data := AdminEditClientPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			UserData:     currentUser,
		}
		err = tmpl.ExecuteTemplate(w, "layout", data)
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
		fmt.Println(user)
		defer file.Close()
		if r.FormValue("speichern") == "2" {
			user.Update()
			fmt.Println("User wurde geupdated")
			fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
			currentUserBearb, _ := benutzer.Get(2)
			data := AdminEditClientPageData{
				Benutzername: currentBenutzerName,
				BenutzerTyp:  currentBenutzerTyp,
				UserData:     currentUserBearb,
			}
			err = tmpl.ExecuteTemplate(w, "layout", data)
		}

		f, err := os.OpenFile("./static/images/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentUserbearb, _ := benutzer.Get(id)
	currentUserbearb.Sperren()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde gesperrt")

	currentBenutzerName := "Peter Dieter"
	currentBenutzerTyp := "Benutzer"
	currentUser, _ := benutzer.Get(id)
	data := AdminEditClientPageData{
		Benutzername: currentBenutzerName,
		BenutzerTyp:  currentBenutzerTyp,
		UserData:     currentUser,
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Println(err)
	}
}
