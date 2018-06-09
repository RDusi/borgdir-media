package bookmarked

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	"github.com/jhoefker/borgdir-media/app/model/categorie"
	"github.com/jhoefker/borgdir-media/app/model/equipment"
	"github.com/jhoefker/borgdir-media/app/model/storage"
)

type BookmarkedItem struct {
	ID             int
	User           benutzer.User
	Equipment      equipment.Equipment
	RueckgabeDatum string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data/borgdir.media.db")
	if err != nil {
		panic(err)
	}
}

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
		var lagerortnr int
		var kategorienr int
		categorie := categorie.Categorie{}
		storage := storage.Storage{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
		err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
		equipment.Lagerort = storage
		equipment.Kategorie = categorie
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
		var lagerortnr int
		var kategorienr int
		categorie := categorie.Categorie{}
		storage := storage.Storage{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
		err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
		equipment.Lagerort = storage
		equipment.Kategorie = categorie
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
		var lagerortnr int
		var kategorienr int
		categorie := categorie.Categorie{}
		storage := storage.Storage{}
		err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
		err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
		equipment.Lagerort = storage
		equipment.Kategorie = categorie
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
	bookmarkeditem = BookmarkedItem{}
	user := benutzer.User{}
	equipment := equipment.Equipment{}
	storage := storage.Storage{}
	categorie := categorie.Categorie{}
	var userid int
	var equipid int
	var kategorienr int
	var lagerortnr int
	err = Db.QueryRow("select * from Vorgemerkt where id = $1", id).Scan(&bookmarkeditem.ID, &userid, &equipid, &bookmarkeditem.RueckgabeDatum)
	err = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
	err = Db.QueryRow("select * from Equipment where id = $1", equipid).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
	err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
	err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
	equipment.Lagerort = storage
	equipment.Kategorie = categorie
	bookmarkeditem.User = user
	bookmarkeditem.Equipment = equipment
	return
}

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
