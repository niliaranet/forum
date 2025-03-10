package repository

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/niliaranet/forum/models"
	"github.com/niliaranet/forum/utils"
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
		content text,
		time timestamp default current_timestamp
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

func GetPost(id string) models.Post {
	var name string
	var content string
	var time time.Time

	_ = database.QueryRow("select name, content, time from post where id = ?;", id).Scan(&id, &name, &content, &time)
	return models.Post{Name: name, Content: content, Time: utils.FormatTimestamp(time)}
}

func GetPosts() []models.Post {
	sqlStmt := `
	select id, name, content, time from post
	order by id desc;
	`

	rows, err := database.Query(sqlStmt)
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var id int
		var name string
		var content string
		var time time.Time
		err = rows.Scan(&id, &name, &content, &time)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, models.Post{
			Id:      id,
			Name:    name,
			Content: content,
			Time:    utils.FormatTimestamp(time),
		})
	}

	return posts
}

func CreatePost(post models.Post) {
	_, err := database.Exec(` 
	insert into post 
		(name, content)
	values
		(?, ?);
	`, post.Name, post.Content)

	if err != nil {
		log.Print(err)
		return
	}
}
