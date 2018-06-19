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
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Verleiher" {
		http.Redirect(w, r, "/admin/index", http.StatusFound)
	} else {
		fmt.Println("CartHandler")
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			// GET
			t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-cart.tmpl", "template/user/cart.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
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
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentCartItem, _ := model.GetWarenkorbItemByID(id)
	currentEquipment, _ := model.GetEquipmentByID(currentCartItem.Equipment.ID)
	currentEquipment.Anzahl++
	currentEquipment.Update()
	currentCartItem.Delete()
	http.Redirect(w, r, "/cart", 301)
}

func RentItems(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	auth := session.Values["authenticated"].(bool)
	if auth == false {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
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
}
