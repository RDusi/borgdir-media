package database

import "database/sql"

func SetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./borgdirmediaDB.db")
	CheckErr(err)
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
