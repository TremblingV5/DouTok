package comment

import (
	"fmt"
	"time"

	"github.com/TremblingV5/DouTok/applications/comment/infra/model"
)

type Entity struct {
	Id             int64
	VideoId        int64
	UserId         int64
	ConversationId int64
	LastId         int64
	ToUserId       int64
	Content        string
	Timestamp      string
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (c *Entity) check(options ...EntityCheckOption) error {
	for _, option := range options {
		if err := option(c); err != nil {
			return err
		}
	}

	return nil
}

func (c *Entity) isValid() error {
	return c.check(
		IsVideoIdValid(),
		IsUserIdValid(),
	)
}

func New(options ...EntityOption) (*Entity, error) {
	comment := &Entity{}

	for _, option := range options {
		option(comment)
	}

	return comment, comment.isValid()
}

func (c *Entity) ToModel() *model.Comment {
	return &model.Comment{
		Id:             c.Id,
		VideoId:        c.VideoId,
		UserId:         c.UserId,
		ConversationId: c.ConversationId,
		LastId:         c.LastId,
		ToUserId:       c.ToUserId,
		Content:        c.Content,
		Timestamp:      c.Timestamp,
		Status:         c.Status,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}
}

func (c *Entity) ToHBModel() *model.CommentInHB {
	return &model.CommentInHB{
		Id:             []byte(fmt.Sprint(c.Id)),
		VideoId:        []byte(fmt.Sprint(c.VideoId)),
		UserId:         []byte(fmt.Sprint(c.UserId)),
		ConversationId: []byte(fmt.Sprint(c.ConversationId)),
		LastId:         []byte(fmt.Sprint(c.LastId)),
		ToUserId:       []byte(fmt.Sprint(c.ToUserId)),
		Content:        []byte(c.Content),
		Timestamp:      []byte(c.Timestamp),
	}
}

func (c *Entity) TransformFromModel(comment *model.Comment) {
	c.Id = comment.Id
	c.VideoId = comment.VideoId
	c.UserId = comment.UserId
	c.ConversationId = comment.ConversationId
	c.LastId = comment.LastId
	c.ToUserId = comment.ToUserId
	c.Content = comment.Content
	c.Timestamp = comment.Timestamp
	c.Status = comment.Status
	c.CreatedAt = comment.CreatedAt
	c.UpdatedAt = comment.UpdatedAt
}

func (c *Entity) MarkAsDelete() {
	c.Status = false
}

func (c *Entity) IsBelongUser(userId int64) (bool, error) {
	err := c.check(IsBelongToUser(userId))
	return err == nil, err
}
