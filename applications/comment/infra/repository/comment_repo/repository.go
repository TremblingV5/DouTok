package comment_repo

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/comment/infra/model"
	"github.com/TremblingV5/DouTok/applications/comment/infra/query"
	"github.com/TremblingV5/box/dbtx"
)

type Repository interface {
	Save(ctx context.Context, comment *model.Comment) error
	Update(ctx context.Context, comment *model.Comment) error
	MarkAsDeleted(ctx context.Context, id int64) error
	LoadById(ctx context.Context, id int64) (*model.Comment, error)
}

type PersistRepository struct {
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (r *PersistRepository) Save(ctx context.Context, comment *model.Comment) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		return tx.Comment.WithContext(ctx).Save(comment)
	})
}

func (r *PersistRepository) Update(ctx context.Context, comment *model.Comment) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := tx.Comment.WithContext(ctx).Where(
			query.Comment.Id.Eq(comment.Id),
		).Updates(comment)

		return err
	})
}

func (r *PersistRepository) MarkAsDeleted(ctx context.Context, id int64) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := tx.Comment.WithContext(ctx).Where(
			query.Comment.Id.Eq(id),
		).Update(
			query.Comment.Status, false,
		)

		return err
	})
}

func (r *PersistRepository) LoadById(ctx context.Context, id int64) (comment *model.Comment, err error) {
	err = dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		c, err := tx.Comment.WithContext(ctx).Where(query.Comment.Id.Eq(id)).First()
		comment = c
		return err
	})

	return comment, err
}
