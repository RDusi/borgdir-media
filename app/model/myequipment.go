package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type MyEquipItem struct {
	ID             int
	User           User
	Equipment      Equipment
	EntleihDatum   string
	RueckgabeDatum string
}

func GetAllMeineGeraete() (myequipitems []MyEquipItem, err error) {
	rows, err := Db.Query("select * from MeineGeraete")

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		myequipitem := MyEquipItem{}
		err = rows.Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
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
		myequipitem.User = user
		myequipitem.Equipment = equipment
		if err != nil {
			return
		}
		myequipitems = append(myequipitems, myequipitem)
	}

	rows.Close()
	return
}

func GetAllMeineGeraeteByUserId(uid int) (myequipitems []MyEquipItem, err error) {
	rows, err := Db.Query("select * from MeineGeraete where UserID = $1", uid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		myequipitem := MyEquipItem{}
		err = rows.Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
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
		myequipitem.User = user
		myequipitem.Equipment = equipment
		if err != nil {
			return
		}
		myequipitems = append(myequipitems, myequipitem)
	}

	rows.Close()
	return
}

func GetAllMeineGeraeteByEquipmentID(eid int) (myequipitems []MyEquipItem, err error) {
	rows, err := Db.Query("select * from MeineGeraete where EquipmentID = $1", eid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		myequipitem := MyEquipItem{}
		err = rows.Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
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
		myequipitem.User = user
		myequipitem.Equipment = equipment
		if err != nil {
			return
		}
		myequipitems = append(myequipitems, myequipitem)
	}

	rows.Close()
	return
}

func GetMeineGeraeteByID(id int) (myequipitem MyEquipItem, err error) {
	myequipitem = MyEquipItem{}
	user := User{}
	equipment := Equipment{}
	storage := Storage{}
	categorie := Categorie{}
	var userid int
	var equipid int
	var kategorienr int
	var lagerortnr int
	err = Db.QueryRow("select * from MeineGeraete where id = $1", id).Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
	err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
	err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
	err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
	err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
	equipment.Lagerort = storage
	equipment.Kategorie = categorie
	myequipitem.User = user
	myequipitem.Equipment = equipment
	return
}

func (myequipitem *MyEquipItem) Add() (err error) {
	statement := "insert into MeineGeraete (UserID, EquipmentID, EntleihDatum, RueckgabeDatum) values ($1, $2, $3, $4)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(myequipitem.User.ID, myequipitem.Equipment.ID, myequipitem.EntleihDatum, myequipitem.RueckgabeDatum)
	return
}

func (myequipitem *MyEquipItem) Delete() (err error) {
	_, err = Db.Exec("delete from MeineGeraete where id = $1", myequipitem.ID)
	return
}

func (myequipitem *MyEquipItem) Update() (err error) {
	statement := "update MeineGeraete set UserID = ?, EquipmentID = ?, EntleihDatum= ?, RueckgabeDatum= ? where id = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(myequipitem.User.ID, myequipitem.Equipment.ID, myequipitem.EntleihDatum, myequipitem.RueckgabeDatum, myequipitem.ID)
	return
}
