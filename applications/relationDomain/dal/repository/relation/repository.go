package relation

import (
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"gorm.io/gorm"
)

type Repository interface {
	CreateOrUpdate(relation *pack.Relation) error
	CreateList(relationList []*pack.Relation) error
}

type PersistRepository struct {
	relation query.IRelationDo
}

func New(db *gorm.DB) *PersistRepository {
	return &PersistRepository{
		relation: query.Relation.WithContext(db.Statement.Context),
	}
}

func packToModel(rel *pack.Relation) *model.Relation {
	return &model.Relation{
		UserId:   rel.UserId,
		ToUserId: rel.ToUserId,
		Status:   int(rel.ActionType),
	}
}

func (p *PersistRepository) CreateOrUpdate(rel *pack.Relation) error {
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

func (p *PersistRepository) CreateList(relations []*pack.Relation) error {
	models := make([]*model.Relation, 0, len(relations))
	for _, relation := range relations {
		models = append(models, packToModel(relation))
	}
	err := p.relation.CreateInBatches(models, len(models))
	if err != nil {
		return err
	}
	return nil
}
