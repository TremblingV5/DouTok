package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func BuildMessageActionResp(err error, resp *messageDomain.DoutokAddMessageResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildMessageChatResp(err error, resp *messageDomain.DoutokListMessageResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}

func BuildMessageFriendResp(err error, resp *message.DouyinFriendListMessageResponse) {
	e := errno.ConvertErr(err)
	resp.StatusMsg = e.ErrMsg
	resp.StatusCode = int32(e.ErrCode)
}
