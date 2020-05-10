package models

import (
	"time"
	"database/sql"
)

// Post -пост для блога
type Post struct {
	ID int
	Name string
	Post string
	CreatedAt time.Time
}

// PostsSlice - array Posts
type PostsSlice []Post


// ShowAll - return all posts
func ShowAll(db *sql.DB) (PostsSlice, error) {
	res := []Post{}
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		myShowTable := Post{}

		err := rows.Scan(&myShowTable.ID, &myShowTable.Name, &myShowTable.Post, &myShowTable.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, myShowTable)
	}
	return res, err
}


// fmt.Println(res)