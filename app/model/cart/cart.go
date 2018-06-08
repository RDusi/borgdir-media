package cart

import (
	"database/sql"
	"time"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type CartItem struct {
	ID             int
	User           benutzer.User
	Equipment      equipment.Equipment
	EntleihDatum   time.Time
	RueckgabeDatum time.Time
	Anzahl         int
}

// Db handle
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data/borgdir.media.db")
	if err != nil {
		panic(err)
	}
}

// GetAll Cart Items
func GetAll() (cartitems []CartItem, err error) {
	rows, err := Db.Query("select * from Warenkorb")

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		cartitem := CartItem{}
		err = rows.Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		cartitem.User = user
		cartitem.Equipment = equipment
		if err != nil {
			return
		}
		cartitems = append(cartitems, cartitem)
	}

	rows.Close()
	return
}

// GetAll Cart Items By UserID
func GetAllByUserId(uid int) (cartitems []CartItem, err error) {
	rows, err := Db.Query("select * from Warenkorb where UserID = $1", uid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		cartitem := CartItem{}
		err = rows.Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		cartitem.User = user
		cartitem.Equipment = equipment
		if err != nil {
			return
		}
		cartitems = append(cartitems, cartitem)
	}

	rows.Close()
	return
}

// GetAll Cart Items By EquipID
func GetAllByEquipmentID(eid int) (cartitems []CartItem, err error) {
	rows, err := Db.Query("select * from Warenkorb where EquipmentID = $1", eid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		cartitem := CartItem{}
		err = rows.Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		cartitem.User = user
		cartitem.Equipment = equipment
		if err != nil {
			return
		}
		cartitems = append(cartitems, cartitem)
	}

	rows.Close()
	return
}

func Get(id int) (cartitem CartItem, err error) {
	rows, err := Db.Query("select * from Warenkorb where ID = $1", id)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		err = rows.Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		cartitem.User = user
		cartitem.Equipment = equipment
		if err != nil {
			return
		}
	}
	rows.Close()
	return
}

// Add CartItem
func (cartitem *CartItem) Add() (err error) {
	statement := "insert into Warenkorb (UserID, EquipmentID, EntleihDatum, RueckgabeDatum, Anzahl) values ($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(cartitem.User.ID, cartitem.Equipment.ID, cartitem.EntleihDatum, cartitem.RueckgabeDatum, cartitem.Anzahl)
	return
}

func (cartitem *CartItem) Delete() (err error) {
	_, err = Db.Exec("delete from Warenkorb where id = $1", cartitem.ID)
	return
}

func (cartitem *CartItem) Update() (err error) {
	statement := "update Warenkorb set UserID = ?, EquipmentID = ?, EntleihDatum= ?, RueckgabeDatum= ?, Anzahl= ? where id = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(cartitem.User.ID, cartitem.Equipment.ID, cartitem.EntleihDatum, cartitem.RueckgabeDatum, cartitem.Anzahl, cartitem.ID)
	return
}
