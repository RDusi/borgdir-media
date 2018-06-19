package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model"
)

type SliderData struct {
	EquipmentListe []model.Equipment
	Startbild      string
}

func renderSliderBilder() SliderData {
	liste, _ := model.GetAllEquipment()
	startdata := SliderData{
		EquipmentListe: liste[1:],
		Startbild:      liste[0].Bild,
	}
	return startdata
}

func IndexStartHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, _ := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/equipment", http.StatusFound)
	} else if user.BenutzerTyp == "Verleiher" {
		http.Redirect(w, r, "/admin/index", http.StatusFound)
	} else {
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/guest/header/header-std.tmpl", "template/guest/index-start.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		data := renderSliderBilder()
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
