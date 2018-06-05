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
var currentUserglob benutzer.User

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
	currentUserglob = currentUser

	if r.Method == "GET" {
		// GET
		data := AdminEditClientPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			UserData:     currentUser,
		}
		if data.UserData.AktivBis == "gesperrt" {
			tmpl, err = template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client-gesperrt.tmpl")
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("current User", currentUserglob.ID)
		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}

	if r.Method == "POST" {
		// POST
		fmt.Println("current User Post", currentUserglob.ID)
		r.ParseForm()
		r.ParseMultipartForm(32 << 20)
		// logic part of Profil

		benutzername := r.FormValue("benutzername")
		email := r.FormValue("email")
		file, handler, err := r.FormFile("uploadfile")
		speichern, _ := strconv.Atoi(r.FormValue("speichern"))
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

		defer file.Close()
		currentBenutzerTyp = "Benutzer"
		aktivBis := currentUser.AktivBis
		user := benutzer.User{ID: speichern, Benutzername: benutzername, Email: email, Passwort: passwort, Bild: bild, BenutzerTyp: currentBenutzerTyp, AktivBis: aktivBis}
		fmt.Println(currentUser.ID)
		fmt.Println(user)
		fmt.Println("Speichern", r.FormValue("speichern"))
		user.Update()
		fmt.Println("User wurde geupdated")
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		currentUserBearb, _ := benutzer.Get(speichern)
		data := AdminEditClientPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			UserData:     currentUserBearb,
		}
		err = tmpl.ExecuteTemplate(w, "layout", data)

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
	tmpl, err = template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client-gesperrt.tmpl")
	if err != nil {
		fmt.Println(err)
	}
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

func DeblockUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentUserbearb, _ := benutzer.Get(id)
	currentUserbearb.Entsperren()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde entsperrt")
	tmpl, err = template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client.tmpl")
	if err != nil {
		fmt.Println(err)
	}
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
