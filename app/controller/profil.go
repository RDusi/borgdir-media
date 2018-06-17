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
		fmt.Println(speichern)

		if err != nil {
			fmt.Println(err)
			return
		}

		bild := "../../../static/images/" + handler.Filename
		fmt.Println(currentUserEdit)
		user, _ := model.GetBenutzerByID(currentUserEdit.ID)
		fmt.Println(user)
		if speichern == user.ID {
			fmt.Println("hier")
			if benutzername != "" && benutzername != user.Benutzername {
				user.Benutzername = benutzername
			}
			if email != "" && email != user.Email {
				user.Email = email
			}
			user.Bild = bild
			passwordDB, _ := base64.StdEncoding.DecodeString(user.Passwort) //PW aus DB
			err1 := bcrypt.CompareHashAndPassword(passwordDB, []byte(passwortalt))
			if err1 == nil {
				hashedPwdNeu, _ := bcrypt.GenerateFromPassword([]byte(passwortneu), 14)
				b64HashedNeuPwd := base64.StdEncoding.EncodeToString(hashedPwdNeu)
				passwordDBNeuDE, _ := base64.StdEncoding.DecodeString(b64HashedNeuPwd) //PW aus DB
				err2 := bcrypt.CompareHashAndPassword(passwordDBNeuDE, []byte(passwortneuwdh))
				if err2 == nil {
					fmt.Println("neues Passwort korrekt")
					user.Passwort = BytesToString(passwordDBNeuDE)
					user.Update()
					http.Redirect(w, r, "/profil", http.StatusFound)
				}
			}
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
