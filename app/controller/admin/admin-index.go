package admin

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IndexAdminHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-index.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		//data := profil.ProfilDummyDataAdmin()

		err = t.ExecuteTemplate(w, "layout", "data")
		if err != nil {
			fmt.Println(err)
		}
	}
}
