package dao

type CommentOperation interface {
	AddComment(comment Comment) error

	DelComment(comment Comment) error

	GetVideoComment(videoId int64) []Comment
}
