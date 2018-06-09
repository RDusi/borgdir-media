package guest

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/bookmarked"
	"github.com/jhoefker/borgdir-media/app/model/cart"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
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
		currentUser := benutzer.User{ID: 0, Benutzername: "Peter Test", BenutzerTyp: "Benutzer"}
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
	fmt.Println("ist hier")
	id, _ := strconv.Atoi(r.FormValue("id"))
	fmt.Println(id)
	currentEquip, _ := equipment.Get(id)
	fmt.Println(currentEquip)
	var cartItem cart.CartItem
	cartItem.Equipment = currentEquip
	cartItem.EntleihDatum = time.Now()
	cartItem.RueckgabeDatum = time.Now().AddDate(0, 2, 0)
	cartItem.Anzahl = 1
	fmt.Println("Current Item: ", cartItem)
	cartItem.Add()
	fmt.Println("wurde hinzugefuegt")
	http.Redirect(w, r, "/equipment", 301)
}

func Bookmark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ist hier")
	id, _ := strconv.Atoi(r.FormValue("id"))
	fmt.Println(id)
	currentEquip, _ := equipment.Get(id)
	fmt.Println(currentEquip)
	var bookmarkItem bookmarked.BookmarkedItem
	bookmarkItem.Equipment = currentEquip
	bookmarkItem.RueckgabeDatum = time.Now().AddDate(0, 2, 0)
	fmt.Println("Current Item: ", bookmarkItem)
	bookmarkItem.Add()
	http.Redirect(w, r, "/equipment", 301)
}
