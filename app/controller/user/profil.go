package user

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
)

type ProfilPageData struct {
	Benutzername string
	BenutzerTyp  string
	UserData     benutzer.User
}

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-profil.tmpl", "template/user/profil.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	currentBenutzerName := "Peter Dieter"
	currentBenutzerTyp := "Benutzer"
	currentUser, _ := benutzer.Get(1)

	if r.Method == "GET" {
		// GET
		data := ProfilPageData{
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
		// logic part of Profil

		r.ParseMultipartForm(32 << 20)

		benutzername := r.FormValue("benutzername")
		email := r.FormValue("email")
		passwortalt := r.FormValue("passwortalt")
		passwortneu := r.FormValue("passwortneu")
		passwortneuwdh := r.FormValue("passwortneuwdh")
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		bild := "../../../static/images/" + handler.Filename

		user, _ := benutzer.Get(currentUser.ID)
		var passwort string
		if passwortalt == user.Passwort && passwortneu == passwortneuwdh {
			fmt.Println("gleiches Passwort wurde eingegeben")
			passwort = r.FormValue("passwortneu")
		}

		if r.FormValue("speichern") == "2" {
			user := benutzer.User{ID: currentUser.ID, Benutzername: benutzername, Email: email, Passwort: passwort, Bild: bild}
			user.Update()
		}

		defer file.Close()
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		data := ProfilPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			UserData:     currentUser,
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
