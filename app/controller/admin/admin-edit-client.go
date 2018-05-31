package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func EditClientAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProfilHandler")
	fmt.Println("method:", r.Method)

	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-edit-client.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	if r.Method == "GET" {
		// GET
		//	data := clients.CreateClientDummy()
		err = t.ExecuteTemplate(w, "layout", "data")
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
		// if speichern == 2 {speichern}, if speichern == 1 {Konto sperren}
		speichern := r.FormValue("speichern")

		fmt.Println("Benutzername: ", benutzername)
		fmt.Println("E-Mail: ", email)
		fmt.Println("Passwort Neu: ", passwortneu)
		fmt.Println("Passwort Neu Wdh: ", passwortneuwdh)
		fmt.Println("Konto sperren: ", speichern)

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println("Bild wurde hochgeladen: ", handler.Filename)
		//data := clients.CreateClientDummy()
		err = t.ExecuteTemplate(w, "layout", "data")
		f, err := os.OpenFile("./static/images/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
