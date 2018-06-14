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
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		id, _ := strconv.Atoi(r.FormValue("id"))
		currentUser := model.GetCurrentSession().User
		currentUserEdit, _ := model.GetBenutzerByID(id)
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
		user, _ := model.GetBenutzerByID(speichern)
		user.Benutzername = benutzername
		user.Email = email
		user.Passwort = passwort
		user.Bild = bild
		user.Update()
		fmt.Println("User wurde geupdated")
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		f, err := os.OpenFile("./static/images/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		wert, _ := io.Copy(f, file)
		if wert != 0 {
			http.Redirect(w, r, "/admin/edit-client", http.StatusFound)
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
