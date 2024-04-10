package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func BuildRelationActionResp(err error, resp *relation.DouyinRelationActionResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationCountResp(err error, resp *relation.DouyinRelationCountResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationFollowListResp(err error, resp *relation.DouyinRelationFollowListResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationFollowerListResp(err error, resp *relation.DouyinRelationFollowerListResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationFriendListResp(err error, resp *relation.DouyinRelationFriendListResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}
