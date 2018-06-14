package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jhoefker/borgdir-media/app/model"
)

type MyEquipmentPageData struct {
	User         model.User
	MeineGeraete []model.MyEquipItem
	Vorgemerkte  []model.BookmarkedItem
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

		currentUser := model.GetCurrentSession().User
		fmt.Println(currentUser.ID)
		meineGeraeteListe, _ := model.GetAllMeineGeraeteByUserId(currentUser.ID)
		meineVorgemerktenListe, _ := model.GetAllVorgemerktByUserId(currentUser.ID)
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
	currentMyEquipItem, _ := model.GetMeineGeraeteByID(id)
	t, _ := time.Parse("02.01.2006", currentMyEquipItem.RueckgabeDatum)
	currentMyEquipItem.RueckgabeDatum = t.AddDate(0, 2, 0).Format("02.01.2006")
	currentMyEquipItem.Update()
	fmt.Println("Ausleihvorgang wurde verlaengert")
	http.Redirect(w, r, "/my-equipment", http.StatusFound)
}

func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentBookmark, _ := model.GetVorgemerktByID(id)
	currentBookmark.Delete()
	fmt.Println("Vermerkung wurde geloescht")
	http.Redirect(w, r, "/my-equipment", http.StatusFound)
}
