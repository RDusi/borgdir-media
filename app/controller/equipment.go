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
}

func EquipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EquipmentHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-equip.tmpl", "template/guest/equipment.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		currentUser := model.GetCurrentSession().User
		equipmentListe, err := model.GetAllEquipment()

		data := EquipmentPageData{
			User:           currentUser,
			EquipmentListe: equipmentListe,
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
	id, _ := strconv.Atoi(r.FormValue("id"))
	log.Println("AddToCart von Produkt, ID: ", id)
	currentEquip, _ := model.GetEquipmentByID(id)
	if currentEquip.Anzahl <= 0 {
		currentEquip.Anzahl = 0
		currentEquip.Update()
		http.Redirect(w, r, "/equipment", http.StatusFound)
	} else {
		var cartItem model.CartItem
		cartItem.User = model.GetCurrentSession().User
		cartItem.Equipment = currentEquip
		cartItem.EntleihDatum = time.Now().Format("02.01.2006")
		cartItem.RueckgabeDatum = time.Now().AddDate(0, 2, 0).Format("02.01.2006")
		cartItem.Anzahl = 1
		cartItem.Add()
		http.Redirect(w, r, "/equipment", http.StatusFound)
	}

}

func Bookmark(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	log.Println("Bookmark von Produkt, ID: ", id)
	currentEquip, _ := model.GetEquipmentByID(id)
	var bookmarkItem model.BookmarkedItem
	bookmarkItem.User = model.GetCurrentSession().User
	bookmarkItem.Equipment = currentEquip
	bookmarkItem.RueckgabeDatum = time.Now().AddDate(0, 2, 0).Format("02.01.2006")
	bookmarkItem.Add()
	fmt.Println(bookmarkItem)
	http.Redirect(w, r, "/equipment", http.StatusFound)
}
