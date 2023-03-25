package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
)

func PackageMessageChatResponse(result *messageDomain.DoutokListMessageResponse, err error) (*message.DouyinMessageChatResponse, error) {
	return &message.DouyinMessageChatResponse{
		StatusCode:  result.StatusCode,
		StatusMsg:   result.StatusMsg,
		MessageList: result.Message,
	}, err
}
