package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Equipment struct {
	ID          int
	Bezeichnung string
	Kategorie   Categorie
	InventarNr  string
	Lagerort    Storage
	Inhalt      string
	Anzahl      int
	Hinweise    string
	Bild        string
	User        []User
}

// GetAll Equipment
func GetAllEquipment() (equipments []Equipment, err error) {
	rows, err := Db.Query("select * from Equipment")

	if err != nil {
		return
	}

	var kategorienr int
	var lagerortnr int
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
		categorie := Categorie{}
		err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
		storage := Storage{}
		err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
		equipment.Kategorie = categorie
		equipment.Lagerort = storage
		if err != nil {
			return
		}

		equipments = append(equipments, equipment)
	}

	rows.Close()
	return
}

func GetEquipmentByID(id int) (equipment Equipment, err error) {
	equipment = Equipment{}
	storage := Storage{}
	categorie := Categorie{}
	var kategorienr int
	var lagerortnr int
	err = Db.QueryRow("select * from Equipment where id = $1", id).Scan(&equipment.ID, &equipment.Bezeichnung, &kategorienr, &equipment.InventarNr, &lagerortnr, &equipment.Inhalt, &equipment.Anzahl, &equipment.Hinweise, &equipment.Bild)
	err = Db.QueryRow("select * from Lagerort where id = $1", lagerortnr).Scan(&storage.ID, &storage.LagerortName)
	err = Db.QueryRow("select * from Kategorie where id = $1", kategorienr).Scan(&categorie.ID, &categorie.KategorieName)
	equipment.Lagerort = storage
	equipment.Kategorie = categorie
	return
}

func (equipment *Equipment) Add() (err error) {
	statement := "insert into Equipment (Bezeichnung, Kategorie, InventarNr, Lagerort, Inhalt, Anzahl, Hinweise, Bild) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(equipment.Bezeichnung, equipment.Kategorie.ID, equipment.InventarNr, equipment.Lagerort.ID, equipment.Inhalt, equipment.Anzahl, equipment.Hinweise, equipment.Bild)
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
	_, err = stmt.Exec(equipment.Bezeichnung, equipment.Kategorie.ID, equipment.InventarNr, equipment.Lagerort.ID, equipment.Inhalt, equipment.Anzahl, equipment.Hinweise, equipment.Bild, equipment.ID)
	return
}
