package comment_count

import (
	"time"

	"github.com/TremblingV5/DouTok/applications/comment/infra/model"
)

type Entity struct {
	Id        int64
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(options ...Option) *Entity {
	e := &Entity{}
	for _, option := range options {
		option(e)
	}

	return e
}

func TransformFromModel(model *model.CommentCount) *Entity {
	return &Entity{
		Id:        model.Id,
		Number:    model.Number,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (c *Entity) ToModel() *model.CommentCount {
	return &model.CommentCount{
		Id:        c.Id,
		Number:    c.Number,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (c *Entity) Check(options ...EntityCheckOption) error {
	for _, option := range options {
		if err := option(c); err != nil {
			return err
		}
	}

	return nil
}

func (c *Entity) UpdateNumber(number int64) {
	c.Number = number
}
