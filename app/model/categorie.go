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
