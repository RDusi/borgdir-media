package bookmarked

import (
	"database/sql"
	"time"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
)

type BookmarkedItem struct {
	ID             int
	User           benutzer.User
	Equipment      equipment.Equipment
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
func GetAll() (bookmarkeditems []BookmarkedItem, err error) {
	rows, err := Db.Query("select * from Vorgemerkt")

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		bookmarkeditem := BookmarkedItem{}
		err = rows.Scan(&bookmarkeditem.ID, &userid, &equipid, &bookmarkeditem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		bookmarkeditem.User = user
		bookmarkeditem.Equipment = equipment
		if err != nil {
			return
		}
		bookmarkeditems = append(bookmarkeditems, bookmarkeditem)
	}

	rows.Close()
	return
}

// GetAll Cart Items By UserID
func GetAllByUserId(uid int) (bookmarkeditems []BookmarkedItem, err error) {
	rows, err := Db.Query("select * from Vorgemerkt where UserID = $1", uid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		bookmarkeditem := BookmarkedItem{}
		err = rows.Scan(&bookmarkeditem.ID, &userid, &equipid, &bookmarkeditem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		bookmarkeditem.User = user
		bookmarkeditem.Equipment = equipment
		if err != nil {
			return
		}
		bookmarkeditems = append(bookmarkeditems, bookmarkeditem)
	}

	rows.Close()
	return
}

// GetAll Cart Items By EquipID
func GetAllByEquipmentID(eid int) (bookmarkeditems []BookmarkedItem, err error) {
	rows, err := Db.Query("select * from Vorgemerkt where EquipmentID = $1", eid)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		bookmarkeditem := BookmarkedItem{}
		err = rows.Scan(&bookmarkeditem.ID, &userid, &equipid, &bookmarkeditem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		bookmarkeditem.User = user
		bookmarkeditem.Equipment = equipment
		if err != nil {
			return
		}
		bookmarkeditems = append(bookmarkeditems, bookmarkeditem)
	}

	rows.Close()
	return
}

func Get(id int) (bookmarkeditem BookmarkedItem, err error) {
	rows, err := Db.Query("select * from Vorgemerkt where ID = $1", id)

	if err != nil {
		return
	}
	var userid int
	var equipid int
	for rows.Next() {
		err = rows.Scan(&bookmarkeditem.ID, &userid, &equipid, &bookmarkeditem.RueckgabeDatum)
		user := benutzer.User{}
		err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
		equipment := equipment.Equipment{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		bookmarkeditem.User = user
		bookmarkeditem.Equipment = equipment
		if err != nil {
			return
		}
	}
	rows.Close()
	return
}

// Add CartItem
func (bookmarkeditem *BookmarkedItem) Add() (err error) {
	statement := "insert into Vorgemerkt (UserID, EquipmentID, RueckgabeDatum) values ($1, $2, $3)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(bookmarkeditem.User.ID, bookmarkeditem.Equipment.ID, bookmarkeditem.RueckgabeDatum)
	return
}

func (bookmarkeditem *BookmarkedItem) Delete() (err error) {
	_, err = Db.Exec("delete from Vorgemerkt where id = $1", bookmarkeditem.ID)
	return
}

func (bookmarkeditem *BookmarkedItem) Update() (err error) {
	statement := "update Vorgemerkt set UserID = ?, EquipmentID = ?, RueckgabeDatum= ? where id = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(bookmarkeditem.User.ID, bookmarkeditem.Equipment.ID, bookmarkeditem.RueckgabeDatum, bookmarkeditem.ID)
	return
}
