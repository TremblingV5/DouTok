package comment_count

import "github.com/TremblingV5/DouTok/applications/comment/infra/model"

type CommentCountList []*Entity

func NewListFromModels(models []*model.CommentCount) CommentCountList {
	var result CommentCountList
	for _, model := range models {
		result = append(result, &Entity{
			Id:        model.Id,
			Number:    model.Number,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}

	return result
}
