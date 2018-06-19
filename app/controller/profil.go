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
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
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

			userEDIT, _ := model.GetBenutzerByID(speichern)
			fmt.Println("CURRENTUSER POOOOOOST: ", userEDIT)

			if passwortneu == passwortneuwdh {
				if passwortneuwdh == "" && passwortalt == "" && passwortneuwdh == "" {
					userEDIT.Benutzername = benutzernameINPUT
					userEDIT.Email = email
					userEDIT.Bild = bild
					userEDIT.Update()
					fmt.Println("KEIN NEUES PASSWORT: ", userEDIT)
					http.Redirect(w, r, "/profil", http.StatusFound)
				} else {
					userEDIT.Benutzername = benutzernameINPUT
					userEDIT.Email = email
					userEDIT.Passwort = passwortneu
					userEDIT.Bild = bild
					userEDIT.Update()
					fmt.Println("CURRENTUSER POOOOOOST NACH UPDATE: ", userEDIT)
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
	tmpl, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-profil.tmpl", "template/user/profil-delete.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "layout", "data")
	if err != nil {
		fmt.Println(err)
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}
