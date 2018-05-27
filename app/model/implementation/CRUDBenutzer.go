package implementation

import (
	"database/sql"
	"fmt"

	"github.com/jhoefker/borgdir-media/app/model/entities"
	_ "github.com/mattn/go-sqlite3"
)

func InsertBenutzer(benutzer entities.Benutzer) int64 {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO user(username, userstate, usertype, email, activeuntil, picture, password) values(?,?,?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(benutzer.Username, benutzer.Userstate, benutzer.Usertype, benutzer.Email, benutzer.ActiveUntil, benutzer.Picture, benutzer.Password)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	db.Close()

	return id
}

func GetAlleBenutzer() entities.ListBenutzer {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	var alleBenutzer entities.ListBenutzer
	i := 0
	for rows.Next() {
		err = rows.Scan(&alleBenutzer.Benutzer[i].ID, &alleBenutzer.Benutzer[i].Username, &alleBenutzer.Benutzer[i].Userstate, &alleBenutzer.Benutzer[i].Usertype, &alleBenutzer.Benutzer[i].Email, &alleBenutzer.Benutzer[i].ActiveUntil, &alleBenutzer.Benutzer[i].Picture, &alleBenutzer.Benutzer[i].Password)
		checkErr(err)
		i++
	}

	rows.Close() //good habit to close
	db.Close()

	return alleBenutzer
}

func UpdateBenutzer(benutzer entities.Benutzer) int64 {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	stmt, err := db.Prepare("update user set username=?, userstate=?,  usertype=?, email=?, activeuntil=?, picture=?, password=? where userid=?")
	checkErr(err)

	res, err := stmt.Exec(benutzer.Username, benutzer.Userstate, benutzer.Usertype, benutzer.Email, benutzer.ActiveUntil, benutzer.Picture, benutzer.Password, benutzer.ID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

	return affect
}

func DeleteBenutzer(userid int) int64 {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	stmt, err := db.Prepare("delete from user where userid=?")
	checkErr(err)

	res, err := stmt.Exec(userid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

	return affect
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
