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

type AdminEditClientPageData struct {
	User     model.User
	UserData model.User
	Gesperrt int
}

func EditClientAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("EditClientHandler")
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
			fmt.Println("USEEEEERID: ", id)
			currentUserEdit, _ = model.GetBenutzerByID(id)
			fmt.Println("CURRENTUSER: ", currentUserEdit)
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

			r.ParseForm()
			r.ParseMultipartForm(32 << 20)
			// logic part of Profil
			benutzernameINPUT := r.FormValue("benutzername")
			email := r.FormValue("email")
			passwortneu := r.FormValue("passwortneu")
			passwortneuwdh := r.FormValue("passwortneuwdh")
			file, handler, err := r.FormFile("uploadfile")
			speichern, _ := strconv.Atoi(r.FormValue("speichern"))
			bild := "../../../static/images/" + handler.Filename

			userEDIT, _ := model.GetBenutzerByID(speichern)
			fmt.Println("CURRENTUSER POOOOOOST: ", userEDIT)

			if passwortneu == passwortneuwdh {
				if passwortneuwdh == "" {
					userEDIT.Benutzername = benutzernameINPUT
					userEDIT.Email = email
					userEDIT.Bild = bild
					userEDIT.Update()
					fmt.Println("KEIN NEUES PASSWORT: ", userEDIT)
					http.Redirect(w, r, "/admin/edit-client?id="+strconv.Itoa(userEDIT.ID)+"", http.StatusFound)
				} else {
					userEDIT.Benutzername = benutzernameINPUT
					userEDIT.Email = email
					userEDIT.Passwort = passwortneu
					userEDIT.Bild = bild
					userEDIT.Update()
					fmt.Println("CURRENTUSER POOOOOOST NACH UPDATE: ", userEDIT)
					http.Redirect(w, r, "/admin/edit-client?id="+strconv.Itoa(userEDIT.ID)+"", http.StatusFound)
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
