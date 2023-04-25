package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func BuildRelationActionResp(err error, resp *relationDomain.DoutokAddRelationResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRmRelationActionResp(err error, resp *relationDomain.DoutokRmRelationResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationCountResp(err error, resp *relationDomain.DoutokCountRelationResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationFollowListResp(err error, resp *relationDomain.DoutokListRelationResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationFollowerListResp(err error, resp *relationDomain.DoutokListRelationResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildRelationFriendListResp(err error, resp *relationDomain.DoutokListRelationResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}
