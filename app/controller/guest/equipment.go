package guest

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/bookmarked"
	"github.com/jhoefker/borgdir-media/app/model/cart"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
	"github.com/jhoefker/borgdir-media/app/model/nutzung"
)

type EquipmentPageData struct {
	User           benutzer.User
	EquipmentListe []equipment.Equipment
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
		currentUser := nutzung.GetCurrent().User
		equipmentListe, err := equipment.GetAll()

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
	currentEquip, _ := equipment.Get(id)
	if currentEquip.Anzahl <= 0 {
		currentEquip.Anzahl = 0
		currentEquip.Update()
		http.Redirect(w, r, "/equipment", http.StatusFound)
	} else {
		var cartItem cart.CartItem
		cartItem.User = nutzung.GetCurrent().User
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
	currentEquip, _ := equipment.Get(id)
	var bookmarkItem bookmarked.BookmarkedItem
	bookmarkItem.User = nutzung.GetCurrent().User
	bookmarkItem.Equipment = currentEquip
	bookmarkItem.RueckgabeDatum = time.Now().AddDate(0, 2, 0).Format("02.01.2006")
	bookmarkItem.Add()
	fmt.Println(bookmarkItem)
	http.Redirect(w, r, "/equipment", http.StatusFound)
}
