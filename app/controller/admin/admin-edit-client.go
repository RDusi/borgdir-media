package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/nutzung"
)

type AdminEditClientPageData struct {
	User     benutzer.User
	UserData benutzer.User
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
		currentUser := nutzung.GetCurrent().User
		currentUserEdit, _ := benutzer.Get(id)

		data := AdminEditClientPageData{
			User:     currentUser,
			UserData: currentUserEdit,
		}
		if data.UserData.AktivBis == "gesperrt" {
			tmpl, err = template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client-gesperrt.tmpl")
			if err != nil {
				fmt.Println(err)
			}
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
		user, _ := benutzer.Get(speichern)
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
	currentUserbearb, _ := benutzer.Get(id)
	currentUserbearb.Sperren()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde gesperrt")
	http.Redirect(w, r, "/admin/edit-client?id="+strconv.Itoa(id)+"", http.StatusFound)
}

func DeblockUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentUserbearb, _ := benutzer.Get(id)
	currentUserbearb.Entsperren()
	fmt.Println("Konto mit ID " + strconv.Itoa(id) + " wurde entsperrt")
	http.Redirect(w, r, "/admin/edit-client?id="+strconv.Itoa(id)+"", http.StatusFound)
}
