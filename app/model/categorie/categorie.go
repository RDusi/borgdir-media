package categorie

import "database/sql"

type Categorie struct {
	ID            int
	KategorieName string
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

func Get(id int) (categorie Categorie, err error) {
	categorie = Categorie{}
	err = Db.QueryRow("select * from Kategorie where id = $1", id).Scan(&categorie.ID, &categorie.KategorieName)
	return
}
