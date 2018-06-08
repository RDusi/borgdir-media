package user

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/cart"
)

type CartPageData struct {
	Benutzername string
	BenutzerTyp  string
	CartItems    []cart.CartItem
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

		currentBenutzerName := "Peter Dieter"
		currentBenutzerTyp := "Benutzer"
		cartItems, _ := cart.GetAllByUserId(0)
		data := CartPageData{
			Benutzername: currentBenutzerName,
			BenutzerTyp:  currentBenutzerTyp,
			CartItems:    cartItems,
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
		http.Redirect(w, r, "/cart", 301)
	}
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentCartItem, _ := cart.Get(id)
	currentCartItem.Delete()
	http.Redirect(w, r, "/cart", 301)
}
