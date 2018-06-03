package benutzer

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID           int
	Benutzername string
	Email        string
	Passwort     string
	BenutzerTyp  string
	AktivBis     string
	Bild         string
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
func GetAll() (users []User, err error) {
	rows, err := Db.Query("select * from User")

	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)

		if err != nil {
			return
		}

		users = append(users, user)
	}

	rows.Close()
	return
}

// Get Todo with the provided id
func Get(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from User where id = $1", id).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
	return
}

// Add Todo
func (user *User) Add() (err error) {
	statement := "insert into User (Benutzername, Email, Passwort, BenutzerTyp, AktivBis, Bild) values ($1, $2, $3, $4, $5, $6)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Benutzername, user.Email, user.Passwort, user.BenutzerTyp, user.AktivBis, user.Bild)
	return
}

func (user *User) Update() (err error) {
	_, err = Db.Exec("update User set Benutzername = $2, Email = $3, Passwort = $4, Bild = $5 where id = $1", user.ID, user.Benutzername, user.Email, user.Passwort, user.Bild)
	return
}

func (user *User) Sperren() (err error) {
	_, err = Db.Exec("update User set AktivBis = 'gesperrt' where id = $1", user.ID)
	return
}

// Delete Todo with the provided id from the list of Todos
func (user *User) Delete() (err error) {
	_, err = Db.Exec("delete from User where id = $1", user.ID)
	return
}
