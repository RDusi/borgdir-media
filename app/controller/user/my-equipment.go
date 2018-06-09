package user

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/bookmarked"
	"github.com/jhoefker/borgdir-media/app/model/myequipment"
	"github.com/jhoefker/borgdir-media/app/model/nutzung"
)

type MyEquipmentPageData struct {
	User         benutzer.User
	MeineGeraete []myequipment.MyEquipItem
	Vorgemerkte  []bookmarked.BookmarkedItem
}

func MyEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MyEquipmentHandler")
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// GET
		t, err := template.ParseFiles("template/layout/layout.tmpl", "template/user/header/header-myequip.tmpl", "template/user/my-equipment.tmpl")
		if err != nil {
			fmt.Println(err)
		}

		currentUser := nutzung.GetCurrent().User
		fmt.Println(currentUser.ID)
		meineGeraeteListe, _ := myequipment.GetAllByUserId(currentUser.ID)
		meineVorgemerktenListe, _ := bookmarked.GetAllByUserId(currentUser.ID)
		data := MyEquipmentPageData{
			User:         currentUser,
			MeineGeraete: meineGeraeteListe,
			Vorgemerkte:  meineVorgemerktenListe,
		}
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

func ExtendMyEquipment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentMyEquipItem, _ := myequipment.Get(id)
	currentMyEquipItem.RueckgabeDatum = currentMyEquipItem.RueckgabeDatum + "2 Wochen"
	currentMyEquipItem.Update()
	fmt.Println("Ausleihvorgang wurde verlaengert")
	http.Redirect(w, r, "/my-equipment", 301)
}

func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentBookmark, _ := bookmarked.Get(id)
	currentBookmark.Delete()
	fmt.Println("Vermerkung wurde geloescht")
	http.Redirect(w, r, "/my-equipment", 301)
}
