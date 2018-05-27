package implementation

import (
	"fmt"

	"github.com/jhoefker/borgdir-media/app/model/database"
	"github.com/jhoefker/borgdir-media/app/model/entities"
	_ "github.com/mattn/go-sqlite3"
)

func InsertBenutzer(benutzer entities.Benutzer) int64 {
	db := database.SetDatabase()

	stmt, err := db.Prepare("INSERT INTO User(username, userstate, usertype, email, activeuntil, picture, password) values(?,?,?,?,?,?,?)")
	database.CheckErr(err)

	res, err := stmt.Exec(benutzer.Username, benutzer.Userstate, benutzer.Usertype, benutzer.Email, benutzer.ActiveUntil, benutzer.Picture, benutzer.Password)
	database.CheckErr(err)

	id, err := res.LastInsertId()
	database.CheckErr(err)

	fmt.Println(id)

	db.Close()

	return id
}

func GetBenutzerUndAdmins() entities.ListBenutzer {
	db := database.SetDatabase()

	rows, err := db.Query("SELECT * FROM User")
	database.CheckErr(err)
	var alleBenutzer entities.ListBenutzer
	i := 0
	for rows.Next() {
		err = rows.Scan(&alleBenutzer.Benutzer[i].ID, &alleBenutzer.Benutzer[i].Username, &alleBenutzer.Benutzer[i].Userstate, &alleBenutzer.Benutzer[i].Usertype, &alleBenutzer.Benutzer[i].Email, &alleBenutzer.Benutzer[i].ActiveUntil, &alleBenutzer.Benutzer[i].Picture, &alleBenutzer.Benutzer[i].Password)
		database.CheckErr(err)
		i++
	}
	rows.Close()
	db.Close()

	return alleBenutzer
}

func GetBenutzer() entities.ListBenutzer {
	db := database.SetDatabase()

	rows, err := db.Query("SELECT * FROM User where usertype=1")
	database.CheckErr(err)
	var alleBenutzer entities.ListBenutzer
	i := 0
	for rows.Next() {
		err = rows.Scan(&alleBenutzer.Benutzer[i].ID, &alleBenutzer.Benutzer[i].Username, &alleBenutzer.Benutzer[i].Userstate, &alleBenutzer.Benutzer[i].Usertype, &alleBenutzer.Benutzer[i].Email, &alleBenutzer.Benutzer[i].ActiveUntil, &alleBenutzer.Benutzer[i].Picture, &alleBenutzer.Benutzer[i].Password)
		database.CheckErr(err)
		i++
	}
	rows.Close()
	db.Close()

	return alleBenutzer
}

func UpdateBenutzer(benutzer entities.Benutzer) int64 {
	db := database.SetDatabase()

	stmt, err := db.Prepare("update User set username=?, userstate=?,  usertype=?, email=?, activeuntil=?, picture=?, password=? where userid=?")
	database.CheckErr(err)

	res, err := stmt.Exec(benutzer.Username, benutzer.Userstate, benutzer.Usertype, benutzer.Email, benutzer.ActiveUntil, benutzer.Picture, benutzer.Password, benutzer.ID)
	database.CheckErr(err)

	affect, err := res.RowsAffected()
	database.CheckErr(err)

	fmt.Println(affect)

	db.Close()

	return affect
}

func DeleteBenutzer(userid int) int64 {
	db := database.SetDatabase()

	stmt, err := db.Prepare("delete from User where userid=?")
	database.CheckErr(err)

	res, err := stmt.Exec(userid)
	database.CheckErr(err)
	affect, err := res.RowsAffected()
	database.CheckErr(err)

	fmt.Println(affect)

	db.Close()

	return affect
}
