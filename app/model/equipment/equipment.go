package equipment

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Equipment struct {
	ID          int
	Bezeichnung string
	Kategorie   string
	InventarNr  string
	Lagerort    string
	Inhalt      string
	Anzahl      int
	Hinweise    string
	Bild        string
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

// GetAll Equipment
func GetAll() (equipments []Equipment, err error) {
	rows, err := Db.Query("select * from Equipment")

	if err != nil {
		return
	}

	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)

		if err != nil {
			return
		}

		equipments = append(equipments, equipment)
	}

	rows.Close()
	return
}

// Get Todo with the provided id
func Get(id int) (equipment Equipment, err error) {
	equipment = Equipment{}
	err = Db.QueryRow("select * from Equipment where id = $1", id).Scan(&equipment.ID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNr, &equipment.Lagerort, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
	return
}

// Add Equipment
func (equipment *Equipment) Add() (err error) {
	statement := "insert into Equipment (Bezeichnung, Kategorie, InventarNr, Lagerort, Inhalt, Anzahl, Hinweise, Bild) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(equipment.Bezeichnung, equipment.Kategorie, equipment.InventarNr, equipment.Lagerort, equipment.Inhalt, equipment.Anzahl, equipment.Hinweise, equipment.Bild)
	return
}

func (equipment *Equipment) Delete() (err error) {
	_, err = Db.Exec("delete from Equipment where id = $1", equipment.ID)
	return
}

func (equipment *Equipment) Update() (err error) {
	statement := "update Equipment set Bezeichnung = ?, Kategorie= ?, InventarNr= ?, Lagerort= ?, Inhalt= ?, Anzahl= ?, Hinweise= ?, Bild= ? where id = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(equipment.Bezeichnung, equipment.Kategorie, equipment.InventarNr, equipment.Lagerort, equipment.Inhalt, equipment.Anzahl, equipment.Hinweise, equipment.Bild, equipment.ID)
	return
}
