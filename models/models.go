package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
//import db connector

type Post struct {
	ID int			`json:"id"`
	Title string	`json:"title"`
	Topic string	`json:"topic"`
	Votes int		`json:"votes"`
}

type PostCollection struct {
	Posts []Post 	`json:"items"`
}

func getPosts(db *sql.DB) PostCollection {
	query := "SELECT * FROM posts"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	result := PostCollection{}

	for rows.Next() {
		post := Post{}
		err1 := rows.Scan(&post.ID, &post.Title, &post.Topic, &post.Votes)

		if err1 != nil {
			panic(err)
		}

		result.Posts = append(result.Posts, post)
	}

	return result
}
