package dbFunc

import (
	"time"

	"social-network/structs"
)

func SetComment(dbName string, content string, username string, postid int) error {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()

	insertSQL := `INSERT INTO comments (content, username, postid, creationtime) VALUES (?, ?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	creationtime := time.Now()
	_, err = statement.Exec(content, username, postid, creationtime)
	if err != nil {
		return err
	}
	return nil
}

func GetCommentsByPostId(dbName string, postid int) ([]structs.Comment, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()

	var comment structs.Comment
	comments := make([]structs.Comment, 0)
	rows, err := forumDB.Query("SELECT id, content, username, creationtime FROM comments WHERE postid = ?", postid)
	if err != nil {
		return comments, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&comment.Id, &comment.Content, &comment.Username, &comment.Creationtime); err {
		case nil:
			comment.Postid = postid
			comments = append(comments, comment)
		default:
			return comments, err
		}
	}
	return comments, nil
}
