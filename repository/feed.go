package repository

import (
	"database/sql"
	"simple-social-media/structs"
	"fmt"
)

func getComments(db *sql.DB, feedID int64) ([]structs.CommentResponse, error) {
	commentsSQL := `
			SELECT id, message, created_at, updated_at, user_id
			FROM comments WHERE feed_id = $1
	`

	rows, err := db.Query(commentsSQL, feedID)
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	var comments []structs.CommentResponse
	for rows.Next() {
			var comment structs.Comments
			err := rows.Scan(&comment.ID, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt, &comment.UserID)
			if err != nil {
					return nil, err
			}

			user, err := getUser(db, comment.UserID)
			if err != nil {
				return nil, err
			}

			commentResponse := structs.CommentResponse{
				ID: comment.ID,
				Message: comment.Message,
				CreatedAt: comment.CreatedAt,
				User: user,
			}
			
			comments = append(comments, commentResponse)
	}

	if err := rows.Err(); err != nil {
			return nil, err
	}

	return comments, nil
}

func getLikes(db *sql.DB, feedID int64) ([]structs.LikeResponse, error) {
	likeSQL := `
		SELECT id, created_at, updated_at, user_id
		FROM likes WHERE feed_id = $1
	`

	rows, err := db.Query(likeSQL, feedID)
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	var likes []structs.LikeResponse
	for rows.Next() {
			var like structs.Likes
			err := rows.Scan(&like.ID, &like.CreatedAt, &like.UpdatedAt, &like.UserID)
			if err != nil {
					return nil, err
			}

			user, err := getUser(db, like.UserID)
			if err != nil {
				return nil, err
			}

			likeResponse := structs.LikeResponse{
				ID: like.ID,
				CreatedAt: like.CreatedAt,
				User: user,
			}

			likes = append(likes, likeResponse)
	}

	if err := rows.Err(); err != nil {
			return nil, err
	}

	return likes, nil
}

func getUser(db *sql.DB ,userID int64) (structs.UserResponse, error) {
	query := `
		SELECT id, name FROM users WHERE id = $1
	`

	var user structs.UserResponse
	
	rows, err := db.Query(query, userID)
	if err != nil {
		return user, err
	}
	
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&user.ID, &user.Name)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("no user found for feed ID %d", userID)
		}
		return user, err
	}

	return user, nil
}

func InsertFeed(db *sql.DB, feedData structs.Feeds) (err error) {
	sql := "INSERT INTO feeds (user_id, message, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"

	_, err = db.Exec(sql, feedData.UserID, feedData.Message)
	if err != nil {
			return err
	}


	return nil
}



func UpdateFeed(db *sql.DB, feed structs.Feeds) (err error) {
	sql := "UPDATE feeds SET message = $1, updated_at = NOW() WHERE id = $2"

	_, err = db.Exec(sql, feed.Message, feed.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFeed(db *sql.DB, feed structs.Feeds) (err error) {
	sql := "DELETE FROM feeds WHERE id = $1"

	errs := db.QueryRow(sql, feed.ID)

	sql = "DELETE FROM feeds WHERE id = $1"
	
	errs = db.QueryRow(sql, feed.ID)

	sql = "DELETE FROM feeds WHERE id = $1"
	
	errs = db.QueryRow(sql, feed.ID)

	return errs.Err()
}

func queryGetDataFeed(feedID int64) (sql string) {
	sql = `
			SELECT
			f.id, f.message, f.created_at, f.user_id,
			COUNT(DISTINCT l.id) AS total_likes, 
			COUNT(DISTINCT c.id) AS total_comments
			FROM
					feeds f
			LEFT JOIN
					likes l ON f.id = l.feed_id
			LEFT JOIN
					comments c ON f.id = c.feed_id
			`

	if feedID != 0 {
		sql = sql + fmt.Sprintf(" WHERE f.id = %d GROUP BY f.id", feedID)
	} else {
		sql = sql + fmt.Sprintf(" GROUP BY f.id")
	}

	return sql
}


func GetAllFeed(db *sql.DB) ([]structs.FeedResponse, error) {
	sql := queryGetDataFeed(0)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.FeedResponse
	for rows.Next() {
		var feed = structs.Feeds{}

		err = rows.Scan(&feed.ID, &feed.Message, &feed.CreatedAt, &feed.UserID, &feed.TotalComments, &feed.TotalLikes)
		if err != nil {
			return nil, err
		}

		user, err := getUser(db, feed.UserID)
		if err != nil {
			return nil, err
		}

		feedResponse := structs.FeedResponse{
			ID:            feed.ID,
			Message:       feed.Message,
			CreatedAt:     feed.CreatedAt,
			TotalComments: feed.TotalComments,
			TotalLikes:    feed.TotalLikes,
			User:          user,
		}

		results = append(results, feedResponse)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}


func GetDetailFeed(db *sql.DB, feedID int64) (result structs.FeedDetailResponse, err error) {
	var feed structs.Feeds

	sql := queryGetDataFeed(feedID)
	rows, err := db.Query(sql)
	if err != nil {
		return structs.FeedDetailResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&feed.ID, &feed.Message, &feed.CreatedAt,&feed.UserID, &feed.TotalLikes, &feed.TotalComments)
		if err != nil {
			return structs.FeedDetailResponse{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return structs.FeedDetailResponse{}, err
	}

	user, err := getUser(db, feed.UserID)
	if err != nil {
		return structs.FeedDetailResponse{}, err
	}

	comments, err := getComments(db, feedID)
	if err != nil {
		return structs.FeedDetailResponse{}, err
	}

	likes, err := getLikes(db, feedID)
	if err != nil {
		return structs.FeedDetailResponse{}, err
	}

	if len(comments) == 0 {
		comments = []structs.CommentResponse{}
	}
	
	if len(likes) == 0 {
		likes = []structs.LikeResponse{}
	}

	feedResponse := structs.FeedDetailResponse{
		ID:            feed.ID,
		Message:       feed.Message,
		CreatedAt:     feed.CreatedAt,
		TotalComments: feed.TotalComments,
		TotalLikes:    feed.TotalLikes,
		Comments:				comments,
		Likes:				likes,
		User:          user,
	}


	return feedResponse, nil
}
