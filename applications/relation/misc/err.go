package misc

import "github.com/TremblingV5/DouTok/pkg/errno"

const (
	RelationRepeatedErrCode   = 60001
	RelationUnfollowedErrCode = 60002
)

var (
	RelationRepeatedErr   = errno.NewErrNo(RelationRepeatedErrCode, "repeated relation error")
	RelationUnfollowedErr = errno.NewErrNo(RelationUnfollowedErrCode, "relation unfollowed user error")
)
