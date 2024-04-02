package service

type Relation struct {
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

func NewRelation(userId int64, toUserId int64, actionType int64) *Relation {
	return &Relation{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	}
}
