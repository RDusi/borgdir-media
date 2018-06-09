package myequipment

import (
	"database/sql"
	"time"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type MyEquipItem struct {
	ID             int
	User           benutzer.User
	Equipment      equipment.Equipment
	EntleihDatum   time.Time
	RueckgabeDatum time.Time
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
func GetAll() (myequipitems []MyEquipItem, err error) {
	rows, err := Db.Query("select * from MeineGeraete")

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		myequipitem := MyEquipItem{}
		err = rows.Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
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

// GetAll Cart Items By UserID
func GetAllByUserId(uid int) (myequipitems []MyEquipItem, err error) {
	rows, err := Db.Query("select * from MeineGeraete where UserID = $1", uid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		myequipitem := MyEquipItem{}
		err = rows.Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
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

// GetAll Cart Items By EquipID
func GetAllByEquipmentID(eid int) (myequipitems []MyEquipItem, err error) {
	rows, err := Db.Query("select * from MeineGeraete where EquipmentID = $1", eid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		myequipitem := MyEquipItem{}
		err = rows.Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
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

func Get(id int) (myequipitem MyEquipItem, err error) {
	rows, err := Db.Query("select * from MeineGeraete where ID = $1", id)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		err = rows.Scan(&myequipitem.ID, &userid, &equipid, &myequipitem.EntleihDatum, &myequipitem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		myequipitem.User = user
		myequipitem.Equipment = equipment
		if err != nil {
			return
		}
	}
	rows.Close()
	return
}

// Add CartItem
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