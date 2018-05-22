package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type Cart struct {
	EuqipmentItem equipment.Equipment
	Rueckgabe     string
}

type CartData struct {
	Benutzername   string
	BenutzerStatus string
	Items          []Cart
}

func renderCartData() CartData {
	cartdata := CartData{
		Benutzername:   "Erica Mustermann",
		BenutzerStatus: "Benutzer",
		Items: []Cart{
			{EuqipmentItem: equipment.MyEquipmentListeDummy().Items[0].EuqipmentItem, Rueckgabe: equipment.MyEquipmentListeDummy().Items[0].Rueckgabe},
			{EuqipmentItem: equipment.MyEquipmentListeDummy().Items[1].EuqipmentItem, Rueckgabe: equipment.MyEquipmentListeDummy().Items[0].Rueckgabe},
			{EuqipmentItem: equipment.MyEquipmentListeDummy().Items[2].EuqipmentItem, Rueckgabe: equipment.MyEquipmentListeDummy().Items[0].Rueckgabe},
			{EuqipmentItem: equipment.MyEquipmentListeDummy().Items[3].EuqipmentItem, Rueckgabe: equipment.MyEquipmentListeDummy().Items[0].Rueckgabe},
		},
	}
	return cartdata
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CartHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-cart.tmpl", "template/user/cart.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		data := renderCartData()
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
