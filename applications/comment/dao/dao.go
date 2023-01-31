package dao

import "gorm.io/gorm"

type CommentOperation interface {
	AddComment(comment Comment) error

	DelComment(comment Comment) error

	GetVideoComment(vedioId int64) []Comment
}

type GormDB struct {
	gorm.DB
}

func (g *GormDB) AddComment(comment Comment) error {
	return g.Create(&comment).Error
}

func (g *GormDB) DelComment(comment Comment) error {
	return g.Delete(&comment).Error
}

func (g *GormDB) GetVideoComment(vedioId int64) []Comment {
	var cList []Comment
	g.Find(&cList, "vedio_id=?", vedioId)
	return cList
}
