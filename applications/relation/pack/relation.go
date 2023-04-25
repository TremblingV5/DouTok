package pack

type Relation struct {
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

// func NewRelation[T relationDomain.DoutokRmRelationRequest | relationDomain.DoutokAddRelationRequest](req *T) *Relation {
// 	relation := Relation{
// 		UserId:   req.UserId,
// 		ToUserId: req.ToUserId,
// 		//ActionType: req.ActionType,
// 	}
// 	return &relation
// }

func NewRelation(userId int64, toUserId int64, actionType int64) *Relation {
	return &Relation{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	}
}
