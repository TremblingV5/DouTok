package relation

import (
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"gorm.io/gorm"
)

type Repository interface {
	Save(relation *pack.Relation) error
	SaveList(relationList []*pack.Relation) error

	//LoadOneByUserId(userId int64) (*pack.Relation, error)
	//LoadOneByToUserId(toUserId int64) (*pack.Relation, error)
	//LoadListByUserId(userId int64) ([]*pack.Relation, error)
	//LoadListByToUserId(toUserId int64) ([]*pack.Relation, error)
	//LoadCountByUserId(userId int64) (int64, error)
	//LoadCountByToUserId(toUserId int64) (int64, error)
}

type PersistRepository struct {
	relation query.IRelationDo
}

func New(db *gorm.DB) *PersistRepository {
	return &PersistRepository{
		relation: query.Relation.WithContext(db.Statement.Context),
	}
}

func (p *PersistRepository) Save(rel *pack.Relation) error {
	res, err := p.relation.Where(
		query.Relation.UserId.Eq(rel.UserId),
		query.Relation.ToUserId.Eq(rel.ToUserId),
	).Find()
	if err != nil {
		return err
	}
	if len(res) > 0 {
		// 已经存在关注关系
		_, err := p.relation.Where(
			query.Relation.UserId.Eq(rel.UserId),
			query.Relation.ToUserId.Eq(rel.ToUserId),
		).Update(
			query.Relation.Status, rel.ActionType,
		)
		if err != nil {
			return err
		}
	} else {
		// 不存在则插入
		id := utils.GetSnowFlakeId().Int64()
		err := p.relation.Create(
			&model.Relation{
				ID:       id,
				UserId:   rel.UserId,
				ToUserId: rel.ToUserId,
				Status:   int(rel.ActionType),
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PersistRepository) SaveList(relations []*pack.Relation) error {
	for _, rel := range relations {
		err := p.Save(rel)
		if err != nil {
			return err
		}
	}
	return nil
}
