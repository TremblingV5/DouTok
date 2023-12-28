package comment

import (
	"context"
	"gorm.io/gorm"

	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/query"
)

type Repository struct {
	commentTable query.ICommentDo
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go -package CommentRepositoryMocks
type IRepository interface {
	Save(ctx context.Context, commentId, videoId, userId, conversationId, lastId, toUserId int64, content string, timestamp string) error
	IsCommentFromUser(ctx context.Context, userId, commentId int64) (*model.Comment, bool)
	Remove(ctx context.Context, commentId int64) error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		commentTable: query.Comment.WithContext(db.Statement.Context),
	}
}

func (r *Repository) Save(ctx context.Context, commentId, videoId, userId, conversationId, lastId, toUserId int64, content string, timestamp string) error {
	return r.commentTable.Create(
		&model.Comment{
			Id:             commentId,
			VideoId:        videoId,
			UserId:         userId,
			ConversationId: conversationId,
			LastId:         lastId,
			ToUserId:       toUserId,
			Content:        content,
			Status:         true,
			Timestamp:      timestamp,
		})
}

func (r *Repository) IsCommentFromUser(ctx context.Context, userId, commentId int64) (*model.Comment, bool) {
	comment, err := r.commentTable.Where(
		query.Comment.Id.Eq(commentId),
	).First()
	if err != nil {
		return nil, false
	}

	if comment.UserId == userId {
		return nil, true
	}

	return comment, false
}

func (r *Repository) Remove(ctx context.Context, commentId int64) error {
	_, err := r.commentTable.Where(query.Comment.Id.Eq(commentId)).Update(query.Comment.Status, false)
	return err
}

var _ IRepository = (*Repository)(nil)
