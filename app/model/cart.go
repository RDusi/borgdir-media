package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type CartItem struct {
	ID             int
	User           User
	Equipment      Equipment
	EntleihDatum   string
	RueckgabeDatum string
	Anzahl         int
}

func GetAllWarekorbItems() (cartitems []CartItem, err error) {
	rows, err := Db.Query("select * from Warenkorb")

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		cartitem := CartItem{}
		err = rows.Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
		user := User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := Equipment{}
		var lagerortnr int
		var kategorienr int
		categorie := Categorie{}
		storage := Storage{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
		err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
		equipment.Lagerort = storage
		equipment.Kategorie = categorie
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

func GetAllWarenkorbItemsByUserId(uid int) (cartitems []CartItem, err error) {
	rows, err := Db.Query("select * from Warenkorb where UserID = $1", uid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		cartitem := CartItem{}
		err = rows.Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
		user := User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := Equipment{}
		var lagerortnr int
		var kategorienr int
		categorie := Categorie{}
		storage := Storage{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
		err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
		equipment.Lagerort = storage
		equipment.Kategorie = categorie
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

func GetAllWarenkorbItemsByEquipmentID(eid int) (cartitems []CartItem, err error) {
	rows, err := Db.Query("select * from Warenkorb where EquipmentID = $1", eid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		cartitem := CartItem{}
		err = rows.Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
		user := User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := Equipment{}
		var lagerortnr int
		var kategorienr int
		categorie := Categorie{}
		storage := Storage{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
		err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
		equipment.Lagerort = storage
		equipment.Kategorie = categorie
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

func GetWarenkorbItem(id int) (cartitem CartItem, err error) {
	cartitem = CartItem{}
	user := User{}
	equipment := Equipment{}
	storage := Storage{}
	categorie := Categorie{}
	var userid int
	var equipid int
	var kategorienr int
	var lagerortnr int
	err = Db.QueryRow("select * from Warenkorb where id = $1", id).Scan(&cartitem.ID, &userid, &equipid, &cartitem.EntleihDatum, &cartitem.RueckgabeDatum, &cartitem.Anzahl)
	err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
	err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
	err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
	err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
	equipment.Lagerort = storage
	equipment.Kategorie = categorie
	cartitem.User = user
	cartitem.Equipment = equipment
	return
}

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

func DeleteFromUser(uid int) (err error) {
	_, err = Db.Exec("delete from Warenkorb where UserID = $1", uid)
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
