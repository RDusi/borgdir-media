package guest

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jhoefker/borgdir-media/app/model/cart"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type EquipmentPageData struct {
	Benutzername   string
	BenutzerTyp    string
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
		currentBenutzerName := "Peter Dieter"
		currentBenutzerTyp := "Benutzer"
		equipmentListe, err := equipment.GetAll()
		data := EquipmentPageData{
			Benutzername:   currentBenutzerName,
			BenutzerTyp:    currentBenutzerTyp,
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
	currentEquip, _ := equipment.Get(id)
	var cartItem cart.CartItem
	cartItem.Equipment = currentEquip
	cartItem.EntleihDatum = time.Now()
	cartItem.RueckgabeDatum = time.Now().AddDate(0, 2, 0)
	cartItem.Anzahl = 1
	cartItem.Add()
	http.Redirect(w, r, "/equipment", 301)
}
