package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Categorie struct {
	ID            int
	KategorieName string
}

func GetKategorieById(id int) (categorie Categorie, err error) {
	categorie = Categorie{}
	err = Db.QueryRow("select * from Kategorie where id = $1", id).Scan(&categorie.ID, &categorie.KategorieName)
	return
}

func GetAllKategorien() (categories []Categorie, err error) {
	rows, err := Db.Query("select * from kategorie")

	if err != nil {
		return
	}

	for rows.Next() {
		categorie := Categorie{}
		err = rows.Scan(&categorie.ID, &categorie.KategorieName)

		if err != nil {
			return
		}

		categories = append(categories, categorie)
	}

	rows.Close()
	return
}
