package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jhoefker/borgdir-media/app/model"
)

type MyEquipmentPageData struct {
	User         model.User
	MeineGeraete []model.MyEquipItem
	Vorgemerkte  []model.BookmarkedItem
	AnzahlinCart int
}

func MyEquipmentHandler(w http.ResponseWriter, r *http.Request) {
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
	} else if user.BenutzerTyp == "Verleiher" {
		http.Redirect(w, r, "/admin/index", http.StatusFound)
	} else {
		fmt.Println("MyEquipmentHandler")
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			// GET
			t, err := template.ParseFiles("template/layout.tmpl", "template/header-myequip.tmpl", "template/my-equipment.tmpl")
			if err != nil {
				fmt.Println(err)
			}

			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			fmt.Println(currentUser.ID)
			meineGeraeteListe, _ := model.GetAllMeineGeraeteByUserId(currentUser.ID)
			meineVorgemerktenListe, _ := model.GetAllVorgemerktByUserId(currentUser.ID)
			equips := session.Values["equip"]
			var equip []int
			if equips != nil {
				equip = equips.([]int)
			}
			cartAnzahl := len(equip)

			data := MyEquipmentPageData{
				User:         currentUser,
				MeineGeraete: meineGeraeteListe,
				Vorgemerkte:  meineVorgemerktenListe,
				AnzahlinCart: cartAnzahl,
			}
			err = t.ExecuteTemplate(w, "layout", data)
			if err != nil {
				fmt.Println(err)
			}
		}

		if r.Method == "POST" {
			// POST
			r.ParseForm()
			// logic part of Equipment
		}
	}
}

func ExtendMyEquipment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentMyEquipItem, _ := model.GetMeineGeraeteByID(id)
	t, _ := time.Parse("02.01.2006", currentMyEquipItem.RueckgabeDatum)
	currentMyEquipItem.RueckgabeDatum = t.AddDate(0, 2, 0).Format("02.01.2006")
	currentMyEquipItem.Update()
	fmt.Println("Ausleihvorgang wurde verlaengert")
	http.Redirect(w, r, "/my-equipment", http.StatusFound)
}

func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentBookmark, _ := model.GetVorgemerktByID(id)
	currentBookmark.Delete()
	fmt.Println("Vermerkung wurde geloescht")
	http.Redirect(w, r, "/my-equipment", http.StatusFound)
}
