package pack

import "github.com/TremblingV5/DouTok/kitex_gen/relation"

type Relation struct {
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

func NewRelation(req *relation.DouyinRelationActionRequest) *Relation {
	relation := Relation{
		UserId:     req.UserId,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	}
	return &relation
}
