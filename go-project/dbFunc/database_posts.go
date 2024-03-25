package dbFunc

import (
	"database/sql"
	"time"

	"social-network/structs"
)

func SetPost(dbName string, title string, content string, user int, category int) error {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	username, err := GetUserById(dbName, user)
	if err != nil {
		return err
	}
	insertSQL := `INSERT INTO POSTS(title, content, user, category, creationtime, username) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	creationtime := time.Now()
	_, err = statement.Exec(title, content, user, category, creationtime, username.Username)
	if err != nil {
		return err
	}
	return nil
}

func GetPostById(dbName string, postid int) (structs.Posts, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var post structs.Posts
	post.Id = postid

	row := forumDB.QueryRow("SELECT title, content, user, category, creationtime, username FROM posts WHERE id = ?", postid)

	switch err := row.Scan(&post.Title, &post.Content, &post.User, &post.Category, &post.Creationtime, &post.Username); err {
	case nil:
		return post, nil
	case sql.ErrNoRows:
		return post, err
	default:
		return post, err
	}

}

func GetPostsByCategory(dbName string, category int) ([]structs.Posts, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var post structs.Posts
	posts := make([]structs.Posts, 0)
	rows, err := forumDB.Query("SELECT id, title, content, user, creationtime, username FROM posts WHERE category = ?", category)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.User, &post.Creationtime, &post.Username); err {
		case nil:
			post.Category = category
			posts = append(posts, post)
		default:
			return posts, err
		}
	}
	return posts, nil
}

func GetAllPosts(dbName string) ([]structs.Posts, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var post structs.Posts
	posts := make([]structs.Posts, 0)
	rows, err := forumDB.Query("SELECT id, title, content, user, creationtime, username, category FROM posts")
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.User, &post.Creationtime, &post.Username, &post.Category); err {
		case nil:
			posts = append(posts, post)
		default:
			return posts, err
		}
	}
	return posts, nil
}
