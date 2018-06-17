package model

import (
	"database/sql"
	"encoding/base64"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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

func GetAllBenutzer() (users []User, err error) {
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

func GetBenutzerByID(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from User where id = $1", id).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
	return
}

/*func (user *User) Add() (err error) {
	statement := "insert into User (Benutzername, Email, Passwort, BenutzerTyp, AktivBis, Bild) values ($1, $2, $3, $4, $5, $6)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Benutzername, user.Email, user.Passwort, user.BenutzerTyp, user.AktivBis, user.Bild)
	return
}*/

//////////
func (user *User) Add() (err error) {
	statement := "insert into User (Benutzername, Email, Passwort, BenutzerTyp, AktivBis, Bild) values ($1, $2, $3, $4, $5, $6)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}
	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Passwort), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	_, err = stmt.Exec(user.Benutzername, user.Email, b64HashedPwd, user.BenutzerTyp, user.AktivBis, user.Bild)
	return
}

////////////////////

func (user *User) Update() (err error) {
	statement := "update User set Benutzername = ?, Email= ?, Passwort= ?, BenutzerTyp= ?, AktivBis= ?, Bild= ? where id = ?"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Benutzername, user.Email, user.Passwort, user.BenutzerTyp, user.AktivBis, user.Bild, user.ID)
	return
}

func (user *User) Sperren() (err error) {
	_, err = Db.Exec("update User set AktivBis = 'gesperrt' where id = $1", user.ID)
	return
}

func (user *User) Entsperren() (err error) {
	_, err = Db.Exec("update User set AktivBis = 'aktiv' where id = $1", user.ID)
	return
}

// Delete Todo with the provided id from the list of Todos
func (user *User) Delete() (err error) {
	_, err = Db.Exec("delete from User where id = $1", user.ID)
	return
}

// GetUserByUsername retrieve User by username
func GetUserByUsername(username string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from User where Benutzername = $1", username).Scan(&user.ID, &user.Benutzername, &user.Email, &user.Passwort, &user.BenutzerTyp, &user.AktivBis, &user.Bild)
	return
}
