package user

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/cart"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
	"github.com/jhoefker/borgdir-media/app/model/myequipment"
	"github.com/jhoefker/borgdir-media/app/model/nutzung"
)

type CartPageData struct {
	User      benutzer.User
	CartItems []cart.CartItem
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
		currentUser := nutzung.GetCurrent().User
		cartItems, _ := cart.GetAllByUserId(currentUser.ID)
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
		cartItem, _ := cart.Get(id)
		cartItem.Anzahl = anzahl
		fmt.Println(cartItem)
		cartItem.Update()
		fmt.Println("Update von CartItem Nr: ", id)
		http.Redirect(w, r, "/cart", 301)
	}
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentCartItem, _ := cart.Get(id)
	currentCartItem.Delete()
	http.Redirect(w, r, "/cart", 301)
}

func RentItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ist hier")
	id, _ := strconv.Atoi(r.FormValue("id"))
	var cartItems []cart.CartItem
	cartItems, _ = cart.GetAllByUserId(id)
	fmt.Println(cartItems)
	for _, cartItem := range cartItems {
		var myequipitem myequipment.MyEquipItem
		myequipitem.User = cartItem.User
		myequipitem.Equipment = cartItem.Equipment
		myequipitem.EntleihDatum = cartItem.EntleihDatum
		myequipitem.RueckgabeDatum = cartItem.RueckgabeDatum
		editEquipment, _ := equipment.Get(cartItem.Equipment.ID)
		editEquipment.Anzahl = editEquipment.Anzahl - cartItem.Anzahl
		editEquipment.Update()
	}
	cart.DeleteFromUser(id)

	http.Redirect(w, r, "/cart", 301)
}
