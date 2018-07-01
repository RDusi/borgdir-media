package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jhoefker/borgdir-media/app/model"
)

type EquipmentPageData struct {
	User           model.User
	EquipmentListe []model.Equipment
	AllCategories  []model.Categorie
	AnzahlinCart   int
}

func EquipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EquipmentHandler")
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout.tmpl", "template/header-equip.tmpl", "template/equipment.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		session, _ := store.Get(r, "session")
		var benutzername string
		if session.Values["username"] != nil {
			benutzername = session.Values["username"].(string)
		} else {
			benutzername = ""
		}

		equips := session.Values["equip"]
		var equip []int
		if equips != nil {
			equip = equips.([]int)
		}
		cartAnzahl := len(equip)

		currentUser, _ := model.GetUserByUsername(benutzername)
		equipmentListe, err := model.GetAllEquipment()
		kategorien, _ := model.GetAllKategorien()
		data := EquipmentPageData{
			User:           currentUser,
			EquipmentListe: equipmentListe,
			AllCategories:  kategorien,
			AnzahlinCart:   cartAnzahl,
		}
		fmt.Println("Equipment: ", data)

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

func AddToCart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, _ := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if user.BenutzerTyp == "Verleiher" {
		http.Redirect(w, r, "/equipment", http.StatusFound)
	} else {
		id, _ := strconv.Atoi(r.FormValue("id"))
		log.Println("AddToCart von Produkt, ID: ", id)
		currentEquip, _ := model.GetEquipmentByID(id)
		currentEquip.Anzahl--
		currentEquip.Update()
		if currentEquip.Anzahl < 0 {
			currentEquip.Anzahl = 0
			currentEquip.Update()
			http.Redirect(w, r, "/equipment", http.StatusFound)
		} else {
			if session.Values["equip"] == nil {
				session.Values["equip"] = []int{}
			}
			session.Values["equip"] = append(session.Values["equip"].([]int), id)
			session.Save(r, w)
			http.Redirect(w, r, "/equipment", http.StatusFound)
			fmt.Println("Hinzufügen von Equipment mit ID: ", id)
		}
	}
}

func Bookmark(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Verleiher" {
		http.Redirect(w, r, "/equipment", http.StatusFound)
	} else {
		id, _ := strconv.Atoi(r.FormValue("id"))
		log.Println("Bookmark von Produkt, ID: ", id)
		currentEquip, _ := model.GetEquipmentByID(id)
		var bookmarkItem model.BookmarkedItem
		bookmarkItem.User = user
		bookmarkItem.Equipment = currentEquip
		bookmarkItem.RueckgabeDatum = time.Now().AddDate(0, 2, 0).Format("02.01.2006")
		bookmarkItem.Add()
		fmt.Println(bookmarkItem)
		http.Redirect(w, r, "/equipment", http.StatusFound)
	}
}
