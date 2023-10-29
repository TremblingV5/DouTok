package commentcnt

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/commentDomain/dal/query"
	"gorm.io/gorm"
)

type Repository struct {
	commentCountTable query.ICommentCountDo
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go -package CommentCountRepositoryMocks
type IRepository interface {
	GetCommentsCount(ctx context.Context, videoIdList ...int64) (map[int64]int64, error)
	UpdateCommentsCount(ctx context.Context, videoIdList ...int64) error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		commentCountTable: query.CommentCount.WithContext(db.Statement.Context),
	}
}

func (r *Repository) GetCommentsCount(ctx context.Context, videoIdList ...int64) (map[int64]int64, error) {
	results, err := r.commentCountTable.Where(query.CommentCount.Id.In(videoIdList...)).Find()
	if err != nil {
		return nil, err
	}

	resultMap := make(map[int64]int64)
	for _, result := range results {
		resultMap[result.Id] = result.Number
	}

	// Set comments count to 0 if not exist
	for _, videoId := range videoIdList {
		if _, ok := resultMap[videoId]; !ok {
			resultMap[videoId] = 0
		}
	}

	return resultMap, nil
}

func (r *Repository) UpdateCommentsCount(ctx context.Context, videoIdList ...int64) error {
	return nil
}

var _ IRepository = (*Repository)(nil)
