package model

type CommentInHB struct {
	Id             int64  `json:"id"`
	VideoId        int64  `json:"video_id"`
	UserId         int64  `json:"user_id"`
	ConversationId int64  `json:"conversation_id"`
	LastId         int64  `json:"last_id"`
	ToUserId       int64  `json:"to_user_id"`
	Content        string `json:"content"`
	Timestamp      string `json:"timestamp"`
}
