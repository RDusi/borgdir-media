package route

import (
	"net/http"

	"github.com/jhoefker/borgdir-media/app/controller/admin"
	"github.com/jhoefker/borgdir-media/app/controller/guest"
	"github.com/jhoefker/borgdir-media/app/controller/user"
)

func MapToController() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", guest.IndexStartHandler)
	http.HandleFunc("/login", guest.LoginHandler)
	http.HandleFunc("/register", guest.RegisterHandler)
	http.HandleFunc("/equipment", guest.EquipmentHandler)
	http.HandleFunc("/cart", user.CartHandler)
	http.HandleFunc("/my-equipment", user.MyEquipmentHandler)
	http.HandleFunc("/profil", user.ProfilHandler)
	http.HandleFunc("/admin/", admin.IndexAdminHandler)
	http.HandleFunc("/admin/equipment", admin.EquipmentAdminHandler)
	http.HandleFunc("/delete-equip", admin.DeleteEquipment)
	http.HandleFunc("/admin/add", admin.AddAdminHandler)
	http.HandleFunc("/admin/clients", admin.ClientsAdminHandler)
	http.HandleFunc("/admin/edit-client", admin.EditClientAdminHandler)
	http.HandleFunc("/admin/konto-sperren", admin.BlockUser)
}
