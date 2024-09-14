package email

import (
	"fmt"

	"github.com/St3plox/Blogchain/business/web/broker"
)

type Email struct {
	Recipient string
	Subject   string
	Body      string
}

func LikeEventToEmail(like broker.LikeEvent) Email {

	likeStr := ""
	if like.IsPositive {
		likeStr = "like"
	} else {
		likeStr = "dislike"
	}

	subject := fmt.Sprintf("New %s on your post", likeStr)
	text := fmt.Sprintf(
		"The user %s left %s under post with id: %d",
		like.UserID,
		likeStr,
		like.PostID,
	)

	return Email{
		Recipient: like.UserEmail,
		Subject:   subject,
		Body:      text,
	}
}
