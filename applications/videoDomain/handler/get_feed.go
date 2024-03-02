package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/applications/videoDomain/pack"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (s *VideoDomainServiceImpl) GetFeed(ctx context.Context, req *videoDomain.DoutokGetFeedRequest) (resp *videoDomain.DoutokGetFeedResponse, err error) {
	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}

	user_id_string := misc.FillUserId(fmt.Sprint(req.UserId))

	// 1. 从Redis中获取Feed列表（通过LPop）
	var list []typedef.VideoInHB
	var ok bool
	list, ok = service.GetFeedCache(ctx, user_id_string, 10)

	// 2. 【视频条数不足】从hbase中从latest_time开始，以24h的周期向前查询，直至条数满足或超过current_time - 14 * 24h
	if !ok {
		listFromHB, err := service.SearchFeedEarlierInHB(req.LatestTime, req.LatestTime-7*86400)
		if err != nil {
			return pack.PackageFeedListResp([]typedef.VideoInHB{}, &misc.SystemErr, req.UserId)
		}

		// 3. 取前10条视频作为本次feed的数据，其余的通过RPush进入投递箱
		err = service.SetFeedCache(ctx, "r", user_id_string, listFromHB...)
		if err != nil {
			return pack.PackageFeedListResp([]typedef.VideoInHB{}, &misc.SystemErr, req.UserId)
		}

		var newListNum int64
		if len(listFromHB) >= 10 {
			newListNum = 10
		} else {
			newListNum = int64(len(listFromHB))
		}
		list, ok = service.GetFeedCache(ctx, user_id_string, newListNum)

		if !ok {
			return pack.PackageFeedListResp([]typedef.VideoInHB{}, &misc.SystemErr, req.UserId)
		}
	}

	// 4. 计算current_time与marked_time的差值是否超过6个小时，如是则进行查询
	current_time := time.Now().Unix()
	marked_time, err := service.GetMarkedTime(ctx, user_id_string)
	if err != nil {
		marked_time = fmt.Sprint(current_time)
	}

	if err != nil {
		marked_time = fmt.Sprint(current_time)
		if err := service.SetMarkedTime(ctx, user_id_string, marked_time); err != nil {
			klog.Info("set marked time error")
		}
	}

	if service.JudgeTimeDiff(current_time, marked_time, 60*60*6) {
		// 时间差值已经超过了6个小时
		laterVideoListInHB, new_marked_time, err := service.SearchFeedLaterInHB(marked_time, fmt.Sprint(current_time))
		if err != nil {
			klog.Info("search feed later in hb error")
		}

		if err := service.SetMarkedTime(ctx, user_id_string, new_marked_time); err != nil {
			klog.Info("set marked time error")
		}

		// 5. 若存在新更新的内容，将结果存入投递箱，根据比例选择RPush或LPush
		if err := service.SetFeedCache(ctx, "r", user_id_string, laterVideoListInHB...); err != nil {
			klog.Info("set feed cache error")
		}
	}

	return pack.PackageFeedListResp(list, &misc.Success, req.UserId)
}
