package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/clients"
)

func ClientsAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ClientsAdminHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-clients.tmpl", "template/admin/admin-clients.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		data := clients.ClientListeDummy()
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
