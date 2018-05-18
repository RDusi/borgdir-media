package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

func MyEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MyEquipmentHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-myequip.tmpl", "template/user/my-equipment.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		data := equipment.MyEquipmentListeDummy()
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
