package controller

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model"
	"golang.org/x/crypto/bcrypt"
)

type ProfilPageData struct {
	User     model.User
	UserData model.User
}

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	var benutzername string
	if session.Values["username"] != nil {
		benutzername = session.Values["username"].(string)
	} else {
		benutzername = ""
	}
	user, err := model.GetUserByUsername(benutzername)
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("ProfilHandler")
		fmt.Println("method:", r.Method)
		var currentUser model.User
		var currentUserEdit model.User
		if r.Method == "GET" {
			// GET
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ = model.GetUserByUsername(benutzername.(string))
			currentUserEdit, _ = model.GetBenutzerByID(currentUser.ID)
			fmt.Println(currentUser)
			tmpl, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-profil.tmpl", "template/user/profil.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			data := ProfilPageData{
				User:     currentUser,
				UserData: currentUserEdit,
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

			benutzernameINPUT := r.FormValue("benutzername")
			email := r.FormValue("email")
			passwortalt := r.FormValue("passwortalt")
			passwortneu := r.FormValue("passwortneu")
			passwortneuwdh := r.FormValue("passwortneuwdh")
			file, handler, err := r.FormFile("uploadfile")
			speichern, _ := strconv.Atoi(r.FormValue("speichern"))
			bild := "../../../static/images/" + handler.Filename

			benutzerEdit, _ := model.GetBenutzerByID(speichern)
			if handler.Filename == "" {
				bild = benutzerEdit.Bild
			}
			fmt.Println("CURRENTUSER POOOOOOST: ", benutzerEdit)

			passwordDB, _ := base64.StdEncoding.DecodeString(benutzerEdit.Passwort)
			errPW := bcrypt.CompareHashAndPassword(passwordDB, []byte(passwortalt))

			if passwortneuwdh == "" && passwortalt == "" && passwortneuwdh == "" {
				benutzerEdit.Benutzername = benutzernameINPUT
				benutzerEdit.Email = email
				benutzerEdit.Bild = bild
				benutzerEdit.UpdateWithoutPassword()
				fmt.Println("Profil mit ID: " + strconv.Itoa(benutzerEdit.ID) + " wurde bearbeitet. OHNE NEUES PASSWORT")
				http.Redirect(w, r, "/profil", http.StatusFound)
			} else {
				if errPW == nil && passwortneu == passwortneuwdh {
					fmt.Println("Altes Passwort ist korrekt und neue Passwörter sind gleich")
					benutzerEdit.Benutzername = benutzernameINPUT
					benutzerEdit.Email = email
					benutzerEdit.Passwort = passwortneu
					benutzerEdit.Bild = bild
					benutzerEdit.Update()
					fmt.Println("CURRENTUSER POOOOOOST NACH UPDATE: ", benutzerEdit)
					http.Redirect(w, r, "/profil", http.StatusFound)
				}
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentUserbearb, _ := model.GetBenutzerByID(id)
	currentUserbearb.Delete()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde gelöscht")
	http.Redirect(w, r, "/", http.StatusFound)
}
