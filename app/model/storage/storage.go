package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	ID           int
	LagerortName string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data/borgdir.media.db")
	if err != nil {
		panic(err)
	}
}

func Get(id int) (storage Storage, err error) {
	storage = Storage{}
	err = Db.QueryRow("select * from Lagerort where id = $1", id).Scan(&storage.ID, &storage.LagerortName)
	return
}
