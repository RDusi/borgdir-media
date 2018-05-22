package guest

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type SliderData struct {
	EquipmentListe []equipment.Equipment
	Startbild      string
}

func renderSliderBilder() SliderData {
	liste := equipment.EuqipmentListeDummy()
	startdata := SliderData{
		EquipmentListe: liste.Items[1:],
		Startbild:      liste.Items[0].Bild,
	}
	return startdata
}

func IndexStartHandler(w http.ResponseWriter, r *http.Request) {
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
