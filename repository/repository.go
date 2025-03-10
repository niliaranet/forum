package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Load() {
	os.Remove("database.db")

	var err error
	database, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Panic(err)
	}

	sqlStmt := `
	create table if not exists post (
		id integer not null primary key,
		name text,
		content text
	);

	insert into post 
		(name, content)
	values
		("Hello go!", "This is a post."),
		("How do I exit Vim?", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	;
	`

	_, err = database.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

type post struct {
	Name    string
	Content string
}

func GetPosts() []post {
	sqlStmt := `
	select name, content from post;
	`

	rows, err := database.Query(sqlStmt)
	defer rows.Close()

	var posts []post

	for rows.Next() {
		var name string
		var content string
		err = rows.Scan(&name, &content)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, post{Name: name, Content: content})
	}

	return posts
}
