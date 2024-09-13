package broker

type LikeEvent struct {
	UserID     string `json:"user_id,omitempty"`
	PostID     int64  `json:"post_id"`
	IsPositive bool   `json:"is_positive"`
	UserEmail  string `json:"user_email"`
}
