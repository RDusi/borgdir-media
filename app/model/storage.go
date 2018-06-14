package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	ID           int
	LagerortName string
}

func GetStorageByID(id int) (storage Storage, err error) {
	storage = Storage{}
	err = Db.QueryRow("select * from Lagerort where id = $1", id).Scan(&storage.ID, &storage.LagerortName)
	return
}
