package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model"
)

type IndexPageData struct {
	SliderData   SliderData
	AnzahlinCart int
}

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
	var benutzername string
	if session.Values["username"] != nil {
		benutzername = session.Values["username"].(string)
	} else {
		benutzername = ""
	}
	equips := session.Values["equip"]
	var equip []int
	if equips != nil {
		equip = equips.([]int)
	}
	cartAnzahl := len(equip)
	user, _ := model.GetUserByUsername(benutzername)
	fmt.Println(user)
	if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/equipment", http.StatusFound)
	} else if user.BenutzerTyp == "Verleiher" {
		http.Redirect(w, r, "/admin/index", http.StatusFound)
	} else {
		t, err := template.ParseFiles("template/layout.tmpl", "template/header-std.tmpl", "template/index-start.tmpl")
		if err != nil {
			fmt.Println(err)
		}
		sliderdata := renderSliderBilder()
		data := IndexPageData{
			SliderData:   sliderdata,
			AnzahlinCart: cartAnzahl,
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
