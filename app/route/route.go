package route

import (
	"net/http"

	"github.com/jhoefker/borgdir-media/app/controller"
)

func MapToController() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", controller.IndexStartHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/equipment", controller.EquipmentHandler)
	http.HandleFunc("/addtocart", controller.AddToCart)
	http.HandleFunc("/bookmark", controller.Bookmark)
	http.HandleFunc("/cart", controller.CartHandler)
	http.HandleFunc("/delete-cart-item", controller.DeleteCartItem)
	http.HandleFunc("/cart/rentitems", controller.RentItems)
	http.HandleFunc("/my-equipment", controller.MyEquipmentHandler)
	http.HandleFunc("/my-equipment/extend", controller.ExtendMyEquipment)
	http.HandleFunc("/my-equipment/deleteBookmark", controller.DeleteBookmark)
	http.HandleFunc("/profil", controller.ProfilHandler)
	http.HandleFunc("/konto-loeschen", controller.DeleteUser)
	http.HandleFunc("/admin/", controller.IndexAdminHandler)
	http.HandleFunc("/admin/equipment", controller.EquipmentAdminHandler)
	http.HandleFunc("/delete-equip", controller.DeleteEquipment)
	http.HandleFunc("/admin/add", controller.AddAdminHandler)
	http.HandleFunc("/admin/clients", controller.ClientsAdminHandler)
	http.HandleFunc("/admin/edit-client", controller.EditClientAdminHandler)
	http.HandleFunc("/admin/edit-equipment", controller.AdminEditEquipmentHandler)
	http.HandleFunc("/admin/konto-sperren", controller.BlockUser)
	http.HandleFunc("/admin/konto-entsperren", controller.DeblockUser)
}
