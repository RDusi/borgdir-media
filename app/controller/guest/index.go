package guest

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexStartHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-std.tmpl", "template/guest/index-start.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	data := "...data..."
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Println(err)
	}
}
