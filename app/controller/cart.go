package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model"
)

type CartPageData struct {
	User      model.User
	CartItems []model.CartItem
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
		currentUser := model.GetCurrentSession().User
		cartItems, _ := model.GetAllWarenkorbItemsByUserId(currentUser.ID)
		data := CartPageData{
			User:      currentUser,
			CartItems: cartItems,
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}

	if r.Method == "POST" {
		// POST
		r.ParseForm()
		anzahl, _ := strconv.Atoi(r.FormValue("anzahl"))
		id, _ := strconv.Atoi(r.FormValue("id"))
		fmt.Println(id)
		fmt.Println(anzahl)
		cartItem, _ := model.GetWarenkorbItemByID(id)
		cartItem.Anzahl = anzahl
		fmt.Println(cartItem)
		cartItem.Update()
		fmt.Println("Update von CartItem Nr: ", id)
		http.Redirect(w, r, "/cart", http.StatusFound)
	}
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentCartItem, _ := model.GetWarenkorbItemByID(id)
	currentCartItem.Delete()
	http.Redirect(w, r, "/cart", 301)
}

func RentItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ist hier")
	id, _ := strconv.Atoi(r.FormValue("id"))
	var cartItems []model.CartItem
	cartItems, _ = model.GetAllWarenkorbItemsByUserId(id)
	fmt.Println(cartItems)
	for _, cartItem := range cartItems {
		var myequipitem model.MyEquipItem
		myequipitem.User = cartItem.User
		myequipitem.Equipment = cartItem.Equipment
		myequipitem.EntleihDatum = cartItem.EntleihDatum
		myequipitem.RueckgabeDatum = cartItem.RueckgabeDatum
		myequipitem.Add()
		editEquipment, _ := model.GetEquipmentByID(cartItem.Equipment.ID)
		editEquipment.Anzahl = editEquipment.Anzahl - cartItem.Anzahl
		editEquipment.Update()
	}
	model.DeleteFromUser(id)

	http.Redirect(w, r, "/cart", http.StatusFound)
}