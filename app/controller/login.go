package controller

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jhoefker/borgdir-media/app/model"

	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("Cookie "))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	var benutzername string
	if session.Values["username"] != nil {
		benutzername = session.Values["username"].(string)
	} else {
		benutzername = ""
	}
	user, _ := model.GetUserByUsername(benutzername)
	fmt.Println(user)
	if user.BenutzerTyp == "Verleiher" {
		http.Redirect(w, r, "/admin/index", http.StatusFound)
	} else if user.BenutzerTyp == "Benutzer" && user.AktivBis != "gesperrt" {
		http.Redirect(w, r, "/equipment", http.StatusFound)
	} else {
		fmt.Println("LoginHandler")
		fmt.Println("method:", r.Method)
		T, err := template.ParseFiles("template/layout.tmpl", "template/header-login.tmpl", "template/login.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		session, _ := store.Get(r, "session")
		if r.Method == "GET" {
			// GET
			err = T.ExecuteTemplate(w, "layout", "data")
			if err != nil {
				fmt.Println(err)
			}
		}
		if r.Method == "POST" {
			// POST
			r.ParseForm()

			benutzername := r.FormValue("benutzername")
			fmt.Println(benutzername)
			passwort := r.FormValue("passwort")

			benutzer, _ := model.GetUserByUsername(benutzername)
			fmt.Println(benutzer)
			passwordDB, _ := base64.StdEncoding.DecodeString(benutzer.Passwort)
			err := bcrypt.CompareHashAndPassword(passwordDB, []byte(passwort))

			if err == nil {
				// Set user as authenticated
				session.Values["authenticated"] = true
				session.Values["username"] = benutzername
				session.Values["type"] = benutzer.BenutzerTyp
				session.Save(r, w)
				if benutzer.BenutzerTyp == "Verleiher" {
					http.Redirect(w, r, "/admin", http.StatusFound)
				} else if benutzer.AktivBis == "gesperrt" {
					http.Redirect(w, r, "/login", http.StatusFound)
				} else {
					http.Redirect(w, r, "/", http.StatusFound)
				}
			} else {
				err = T.ExecuteTemplate(w, "layout", "data")
			}

		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	emptyArray := make([]int, 0)
	session, _ := store.Get(r, "session")
	fmt.Println("-----------------------------------")
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Values["equip"] = emptyArray
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
