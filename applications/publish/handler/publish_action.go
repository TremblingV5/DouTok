package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/publish/service"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
)

func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	if ok, msg := check(req); ok {
		resp, _ := service.PackPublishActionRes(1, msg)
		return resp, nil
	}

	if err := service.SavePublish(req.UserId, req.Title, req.Data); err != nil {
		resp, _ := service.PackPublishActionRes(1, err.Error())
		return resp, nil
	}

	resp, _ = service.PackPublishActionRes(0, err.Error())
	return resp, nil
}

func check(req *publish.DouyinPublishActionRequest) (bool, string) {
	if len(req.Data) == 0 {
		return true, "缺少视频数据"
	}

	if len(req.Title) == 0 {
		return true, "缺少标题"
	}

	return false, ""
}
