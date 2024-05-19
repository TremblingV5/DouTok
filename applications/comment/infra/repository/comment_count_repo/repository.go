package comment_count_repo

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/comment/infra/model"
	"github.com/TremblingV5/DouTok/applications/comment/infra/query"
	"github.com/TremblingV5/box/dbtx"
)

type Repository interface {
	Save(ctx context.Context, commentCount *model.CommentCount) error
	Update(ctx context.Context, commentCount *model.CommentCount) error
	LoadById(ctx context.Context, id int64) (*model.CommentCount, error)
	LoadByIdList(ctx context.Context, ids ...int64) ([]*model.CommentCount, error)
	UpdateNumber(ctx context.Context, id, number int64) error
}

type PersistRepository struct{}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (r *PersistRepository) Save(ctx context.Context, commentCount *model.CommentCount) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		return tx.CommentCount.WithContext(ctx).Save(commentCount)
	})
}

func (r *PersistRepository) Update(ctx context.Context, commentCount *model.CommentCount) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := query.CommentCount.WithContext(ctx).Where(
			query.CommentCount.Id.Eq(commentCount.Id),
		).Updates(
			commentCount,
		)

		return err
	})

}

func (r *PersistRepository) LoadById(ctx context.Context, id int64) (commentCount *model.CommentCount, err error) {
	err = dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		c, err := tx.CommentCount.WithContext(ctx).Where(
			query.CommentCount.Id.Eq(id),
		).First()
		commentCount = c
		return err
	})

	return commentCount, err
}

func (r *PersistRepository) LoadByIdList(ctx context.Context, ids ...int64) (commentCountList []*model.CommentCount, err error) {
	err = dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		c, err := tx.CommentCount.WithContext(ctx).Where(
			query.CommentCount.Id.In(ids...),
		).Find()

		commentCountList = append(commentCountList, c...)
		return err
	})

	return commentCountList, err
}

func (r *PersistRepository) UpdateNumber(ctx context.Context, id, number int64) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := query.CommentCount.WithContext(ctx).Where(
			query.CommentCount.Id.Eq(id),
		).Update(
			query.CommentCount.Number, number,
		)
		return err
	})
}
