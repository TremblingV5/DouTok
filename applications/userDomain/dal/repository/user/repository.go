package user

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/query"
	"github.com/TremblingV5/DouTok/pkg/dtdb"
)

type Repository interface {
	Save(ctx context.Context, user *model.User) error
	LoadByUsername(ctx context.Context, username string) (*model.User, error)
	LoadById(ctx context.Context, id uint64) (*model.User, error)
	LoadUserListByIds(ctx context.Context, ids ...uint64) ([]*model.User, error)
	IsUserNameExisted(ctx context.Context, username string) (bool, error)
}

type PersistRepository struct {
	user query.IUserDo
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (p *PersistRepository) Save(ctx context.Context, user *model.User) error {
	tx := dtdb.Tx(ctx).(*query.QueryTx)
	return tx.User.Save(user)
}

func (p *PersistRepository) LoadByUsername(ctx context.Context, username string) (*model.User, error) {
	tx := dtdb.Tx(ctx).(*query.QueryTx)
	return tx.User.Where(query.User.UserName.Eq(username)).First()
}

func (p *PersistRepository) LoadById(ctx context.Context, id uint64) (*model.User, error) {
	tx := dtdb.Tx(ctx).(*query.QueryTx)
	return tx.User.Where(query.User.ID.Eq(id)).First()
}

func (p *PersistRepository) LoadUserListByIds(ctx context.Context, ids ...uint64) ([]*model.User, error) {
	tx := dtdb.Tx(ctx).(*query.QueryTx)
	return tx.User.Where(query.User.ID.In(ids...)).Find()
}

func (p *PersistRepository) IsUserNameExisted(ctx context.Context, username string) (bool, error) {
	tx := dtdb.Tx(ctx).(*query.QueryTx)

	count, err := tx.User.Where(query.User.UserName.Eq(username)).Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
