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

type AdminEditClientPageData struct {
	User     model.User
	UserData model.User
	Gesperrt int
}

func EditClientAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	typ := session.Values["type"]
	if typ.(string) != "Verleiher" && typ == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		fmt.Println("ProfilHandler")
		fmt.Println("method:", r.Method)

		var currentUserEdit model.User

		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client.tmpl")
			if err != nil {
				fmt.Println(err)
			}

			id, _ := strconv.Atoi(r.FormValue("id"))
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			currentUserEdit, _ = model.GetBenutzerByID(id)
			var temp int
			if currentUserEdit.AktivBis == "gesperrt" {
				temp = 1
			}
			data := AdminEditClientPageData{
				User:     currentUser,
				UserData: currentUserEdit,
				Gesperrt: temp,
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
			passwortneu := r.FormValue("passwortneu")
			passwortneuwdh := r.FormValue("passwortneuwdh")
			file, handler, err := r.FormFile("uploadfile")
			speichern, _ := strconv.Atoi(r.FormValue("speichern"))
			if err != nil {
				fmt.Println(err)
				return
			}
			bild := "../../../static/images/" + handler.Filename
			useredit, _ := model.GetBenutzerByID(currentUserEdit.ID)
			if speichern == useredit.ID {
				if benutzername != "" && benutzername != useredit.Benutzername {
					useredit.Benutzername = benutzername
				}
				if email != "" && email != useredit.Email {
					useredit.Email = email
				}
				useredit.Bild = bild
				hashedPwdNeu, _ := bcrypt.GenerateFromPassword([]byte(passwortneu), 14)
				b64HashedNeuPwd := base64.StdEncoding.EncodeToString(hashedPwdNeu)
				passwordDBNeuDE, _ := base64.StdEncoding.DecodeString(b64HashedNeuPwd) //PW aus DB
				err2 := bcrypt.CompareHashAndPassword(passwordDBNeuDE, []byte(passwortneuwdh))
				if err2 == nil {
					fmt.Println("neues Passwort korrekt")
					useredit.Passwort = BytesToString(passwordDBNeuDE)
					useredit.Update()
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
				http.Redirect(w, r, "/profil", http.StatusFound)
			}
		}
	}
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentUserbearb, _ := model.GetBenutzerByID(id)
	currentUserbearb.Sperren()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde gesperrt")
	http.Redirect(w, r, "/admin/edit-client?id="+strconv.Itoa(id)+"", http.StatusFound)
}

func DeblockUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentUserbearb, _ := model.GetBenutzerByID(id)
	currentUserbearb.Entsperren()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde entsperrt")
	http.Redirect(w, r, "/admin/edit-client?id="+strconv.Itoa(id)+"", http.StatusFound)
}
