package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

func EquipmentAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EquipmentAdminHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-equipment.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		data := equipment.EquipmentAdminListeDummy()
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
