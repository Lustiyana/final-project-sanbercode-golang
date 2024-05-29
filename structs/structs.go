package structs

type Users struct {
	ID 					int64		`json:"id"`
	Name 				string	`json:"name"`
	Email			 	string 	`json:"_"`
	Password 		string 	`json:"_"`
}

type Feeds struct {
	ID 						int64		`json:"id"`
	UserID 				int64		`json:"user_id"`
	Message			 	string 	`json:"message"`
	CreatedAt 		string 	`json:"created_at"`
	UpdatedAt 		string 	`json:"updated_at"`
	TotalComments int64		`json:"total_comments"`
	TotalLikes 		int64		`json:"total_likes"`
	Comments []Comments		`json:"_"`
	Likes []Likes		`json:"_"`
	User Users		`json:"user"`
}

type Comments struct {
	ID 						int64		`json:"id"`
	UserID 				int64	`json:"user_id"`
	FeedID			 	int64 	`json:"feed_id"`
	Message				string	`json:"message"`
	CreatedAt 		string 	`json:"created_at"`
	UpdatedAt 		string 	`json:"updated_at"`
	User 		Users 	`json:"user"`
}

type Likes struct {
	ID 						int64		`json:"id"`
	UserID 				int64	`json:"user_id"`
	FeedID			 	int64 	`json:"feed_id"`
	CreatedAt 		string 	`json:"created_at"`
	UpdatedAt 		string 	`json:"updated_at"`
}

type FeedResponse struct {
	ID 						int64		`json:"id"`
	Message			 	string 	`json:"message"`
	CreatedAt 		string 	`json:"created_at"`
	TotalComments int64		`json:"total_comments"`
	TotalLikes 		int64		`json:"total_likes"`
	User 					Users		`json:"user"`
}

type FeedDetailResponse struct {
	ID 						int64		`json:"id"`
	Message			 	string 	`json:"message"`
	CreatedAt 		string 	`json:"created_at"`
	TotalComments int64		`json:"total_comments"`
	TotalLikes 		int64		`json:"total_likes"`
	Comments []CommentResponse		`json:"comments"`
	Likes []LikeResponse		`json:"likes,omitempty"`
	User Users		`json:"user"`
}

type CommentResponse struct {
	ID 						int64		`json:"id"`
	Message				string	`json:"message"`
	CreatedAt 		string 	`json:"created_at"`
	User 		Users 	`json:"user"`
}

type LikeResponse struct {
	ID 						int64		`json:"id"`
	CreatedAt 		string 	`json:"created_at"`
	User Users `json:"user"`
}

type Response struct {
	Success bool       `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}