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

type ProfilPageData struct {
	User     model.User
	UserData model.User
}

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	currentUser := model.GetCurrentSession().User
	if r.Method == "GET" {
		// GET
		tmpl, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-profil.tmpl", "template/user/profil.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		data := ProfilPageData{
			User:     currentUser,
			UserData: currentUser,
		}
		err = tmpl.ExecuteTemplate(w, "layout", data)
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
		speichern, _ := strconv.Atoi(r.FormValue("speichern"))

		if err != nil {
			fmt.Println(err)
			return
		}

		bild := "../../../static/images/" + handler.Filename

		user, _ := model.GetBenutzerByID(currentUser.ID)
		var passwort string
		if passwortalt == user.Passwort && passwortneu == passwortneuwdh {
			fmt.Println("gleiches Passwort wurde eingegeben")
			passwort = r.FormValue("passwortneu")
		}

		userNEW := model.User{ID: speichern, Benutzername: benutzername, Email: email, Passwort: passwort, Bild: bild}
		userNEW.Update()

		defer file.Close()
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		f, err := os.OpenFile("./static/images/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentUserbearb, _ := model.GetBenutzerByID(id)
	currentUserbearb.Delete()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde gelöscht")
	tmpl, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-profil.tmpl", "template/user/profil-delete.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "layout", "data")
	if err != nil {
		fmt.Println(err)
	}
}
