package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database sql.Conn

func Load() {
	os.Remove("env/site.db")

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	sqlStmt := `
	create table if not exists post (
		id integer not null primary key,
		name text,
		content text
	);

	insert into post (
		name, content
	) values (
		"hello go!", "this is a post"
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
