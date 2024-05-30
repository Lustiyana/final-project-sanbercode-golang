package repository

import (
	"database/sql"
	"simple-social-media/structs"
)

func InsertComment(db *sql.DB, commentData structs.Comments) (err error) {
	sql := "INSERT INTO comments (user_id, feed_id, message, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())"

	_, err = db.Exec(sql, commentData.UserID, commentData.FeedID, commentData.Message)
	if err != nil {
			return  err
	}


	return nil
}

func UpdateComment(db *sql.DB, comment structs.Comments) (err error) {
	sql := "UPDATE comments SET message = $1, updated_at = NOW() WHERE id = $2"

	_, err = db.Exec(sql, comment.Message, comment.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteComment(db *sql.DB, comment structs.Comments) (err error) {
	sql := "DELETE FROM comments WHERE id = $1"

	errs := db.QueryRow(sql, comment.ID)

	return errs.Err()
}