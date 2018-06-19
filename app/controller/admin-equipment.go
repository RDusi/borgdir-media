package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jhoefker/borgdir-media/app/model"
)

type ModEquipment struct {
	Equipment model.Equipment
	User      []model.User
	Rueckgabe []string
}

type AdminEquipmentPageData struct {
	User     model.User
	ModEquip []ModEquipment
}

func EquipmentAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, err := model.GetUserByUsername(session.Values["username"].(string))
	fmt.Println(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if user.BenutzerTyp == "Benutzer" {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fmt.Println("EquipmentAdminHandler")
		fmt.Println("method:", r.Method)

		if r.Method == "GET" {
			t, err := template.ParseFiles("template/layout/layout.tmpl", "template/admin/header/header-admin-std.tmpl", "template/admin/admin-equipment.tmpl")
			if err != nil {
				fmt.Println(err)
			}
			session, _ := store.Get(r, "session")
			benutzername := session.Values["username"]
			fmt.Println(benutzername)
			currentUser, _ := model.GetUserByUsername(benutzername.(string))
			currentEquipliste, _ := model.GetAllEquipment()
			var currentEquiplisteTEMP []ModEquipment
			for _, item := range currentEquipliste {
				alleGeraete, _ := model.GetAllMeineGeraeteByEquipmentID(item.ID)
				var alleUser []model.User
				var rueckgaben []string
				for _, geraet := range alleGeraete {
					nutzer, _ := model.GetBenutzerByID(geraet.User.ID)
					rueckgabe := geraet.RueckgabeDatum
					rueckgaben = append(rueckgaben, rueckgabe)
					alleUser = append(alleUser, nutzer)
				}
				item.User = alleUser
				modqeuip := ModEquipment{}
				modqeuip.Equipment = item
				modqeuip.User = alleUser
				modqeuip.Rueckgabe = rueckgaben
				currentEquiplisteTEMP = append(currentEquiplisteTEMP, modqeuip)
			}

			data := AdminEquipmentPageData{
				User:     currentUser,
				ModEquip: currentEquiplisteTEMP,
			}

			fmt.Println(data)
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
}

func DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	currentEquip, _ := model.GetEquipmentByID(id)
	currentEquip.Delete()
	http.Redirect(w, r, "admin/equipment", http.StatusFound)
}
