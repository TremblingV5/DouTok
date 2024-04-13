package comment

import (
	"github.com/TremblingV5/DouTok/applications/comment/infra/model"
)

type List []*Entity

func NewListFromHBModels(models []*model.CommentInHB) List {
	var result List
	for _, model := range models {
		result = append(result, &Entity{
			Id:             model.GetId(),
			VideoId:        model.GetVideoId(),
			UserId:         model.GetUserId(),
			ConversationId: model.GetConversationId(),
			LastId:         model.GetLastId(),
			ToUserId:       model.GetToUserId(),
			Content:        model.GetContent(),
			Timestamp:      model.GetTimestamp(),
		})
	}

	return result
}
