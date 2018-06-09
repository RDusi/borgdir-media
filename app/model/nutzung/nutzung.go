package nutzung

import (
	"database/sql"

	"github.com/jhoefker/borgdir-media/app/model/benutzer"
	_ "github.com/mattn/go-sqlite3"
)

type Session struct {
	ID   int
	User benutzer.User
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

func GetCurrent() (session Session) {
	session = Session{}
	var userid int
	_ = Db.QueryRow("select * from Session").Scan(&session.ID, &userid)
	user := benutzer.User{}
	_ = Db.QueryRow("select * from User where id = $1", userid).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
	session.User = user
	return
}

func (session *Session) Add() (err error) {
	_, err = Db.Exec("delete from Session")
	statement := "insert into Session (UserID) values ($1)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(session.User.ID)
	return
}
