package user

import (
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/query"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user *model.User) error
	LoadByUsername(username string) (*model.User, error)
	LoadById(id uint64) (*model.User, error)
	LoadUserListByIds(ids ...uint64) ([]*model.User, error)
	IsUserNameExisted(username string) (bool, error)
}

type PersistRepository struct {
	user query.IUserDo
}

func New(db *gorm.DB) *PersistRepository {
	return &PersistRepository{
		user: query.User.WithContext(db.Statement.Context),
	}
}

func (p *PersistRepository) Save(user *model.User) error {
	return p.user.Create(user)
}

func (p *PersistRepository) LoadByUsername(username string) (*model.User, error) {
	return p.user.Where(query.User.UserName.Eq(username)).First()
}

func (p *PersistRepository) LoadById(id uint64) (*model.User, error) {
	return p.user.Where(query.User.ID.Eq(id)).First()
}

func (p *PersistRepository) LoadUserListByIds(ids ...uint64) ([]*model.User, error) {
	return p.user.Where(query.User.ID.In(ids...)).Find()
}

func (p *PersistRepository) IsUserNameExisted(username string) (bool, error) {
	count, err := p.user.Where(query.User.UserName.Eq(username)).Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
