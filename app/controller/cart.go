package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jhoefker/borgdir-media/app/model"
)

type CartPageData struct {
	User      model.User
	CartItems []model.CartItem
}

var Items []model.CartItem

func CartHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	var benutzername string
	if session.Values["username"] != nil {
		benutzername = session.Values["username"].(string)
	} else {
		benutzername = ""
	}
	user, _ := model.GetUserByUsername(benutzername)
	fmt.Println(user)
	if user.BenutzerTyp == "Verleiher" {
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
			var benutzername string
			if session.Values["username"] != nil {
				benutzername = session.Values["username"].(string)
			} else {
				benutzername = ""
			}
			currentUser, _ := model.GetUserByUsername(benutzername)
			equips := session.Values["equip"]
			var equip []int
			if equips != nil {
				equip = equips.([]int)
			}
			fmt.Println("Equipment aus Session: ", equip)
			var items []model.CartItem
			for i := 0; i < len(equip); i++ {
				var cartitem model.CartItem
				cartitem.User = currentUser
				cartitem.Equipment, _ = model.GetEquipmentByID(equip[i])
				cartitem.Anzahl = 1
				cartitem.EntleihDatum = time.Now().Format("02.01.2006")
				cartitem.RueckgabeDatum = time.Now().AddDate(0, 2, 0).Format("02.01.2006")
				items = append(items, cartitem)
			}
			Items = items

			data := CartPageData{
				User:      currentUser,
				CartItems: Items,
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
	session, _ := store.Get(r, "session")
	fmt.Println("Delete VOR: Aktuelle Equips in Session: ", session.Values["equip"].([]int))
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentEquipment, _ := model.GetEquipmentByID(id)
	currentEquipment.Anzahl++
	currentEquipment.Update()
	equips := session.Values["equip"]
	var equip []int
	if equips != nil {
		equip = equips.([]int)
	}
	var index int
	for i := 0; i < len(equip); i++ {
		if equip[i] == id {
			index = i
		}
	}
	equipEDIT := append(equip[:index], equip[index+1:]...)
	session.Values["equip"] = equipEDIT
	session.Save(r, w)
	http.Redirect(w, r, "/cart", http.StatusFound)
}

func RentItems(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	auth := session.Values["authenticated"].(bool)
	if auth == false {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("ist hier")
		fmt.Println(Items)
		for _, item := range Items {
			var myequipitem model.MyEquipItem
			myequipitem.User = item.User
			myequipitem.Equipment = item.Equipment
			myequipitem.EntleihDatum = item.EntleihDatum
			myequipitem.RueckgabeDatum = item.RueckgabeDatum
			myequipitem.Add()
		}
		emptyArray := make([]int, 0)
		session.Values["equip"] = emptyArray
		session.Save(r, w)
		http.Redirect(w, r, "/cart", http.StatusFound)
	}
}
