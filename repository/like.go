package repository

import (
	"database/sql"
	"simple-social-media/structs"
)


func InsertLike(db *sql.DB, like structs.Likes) (string, error) {
	checkLikeExist := "SELECT ID FROM likes WHERE user_id = $1 AND feed_id = $2"
	var id int64
	err := db.QueryRow(checkLikeExist, like.UserID, like.FeedID).Scan(&id)
	if err != nil {
			if err == sql.ErrNoRows {
					sql := "INSERT INTO likes (user_id, feed_id, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"
					_, err := db.Exec(sql, like.UserID, like.FeedID)
					if err != nil {
							return "Gagal menyukai", err
					}
					return "Berhasil menyukai", nil
			}
			return "Gagal menyukai", err
	}

	sql := "DELETE FROM likes WHERE user_id = $1 AND feed_id = $2"
	_, err = db.Exec(sql, like.UserID, like.FeedID)
	if err != nil {
			return "Gagal menyukai", err
	}

	return "Batal menyukai", nil
}
