package service

import (
	"fmt"
	"strconv"

	"github.com/TremblingV5/DouTok/applications/publish/misc"
	tools "github.com/TremblingV5/DouTok/pkg/misc"
)

/*
	在HBase中搜索start_time < time <= end_time的视频，作为feed使用
	函数中从HBase中获取了Map结构的数据，并打包成结构体列表
*/
func FindFeedInHB(start_time string, end_time string) ([]VideoInHB, error) {
	start_time_int, _ := strconv.Atoi(start_time)
	end_time_int, _ := strconv.Atoi(end_time)

	res, err := HBClient.ScanRange("feed", misc.GetTimeRebound(int64(end_time_int)), misc.GetTimeRebound(int64(start_time_int)))
	if err != nil {
		return []VideoInHB{}, err
	}

	video_list := []VideoInHB{}
	for _, v := range res {
		temp := VideoInHB{}
		err := tools.Map2Struct4HB(v, &temp)
		if err != nil {
			continue
		}
		video_list = append(video_list, temp)
	}

	return video_list, nil
}

/*
	向前搜索Feed List，前为更早的时间点
*/
func SearchFeedEarlierInHB(latest_time int64, stop_time int64) ([]VideoInHB, error) {
	next_time := latest_time - 86400

	video_list := []VideoInHB{}

	for {
		temp, err := FindFeedInHB(fmt.Sprint(next_time), fmt.Sprint(latest_time))

		if err != nil {
			return video_list, err
		}

		video_list = append(video_list, temp...)

		// 终止条件1：视频列表长度已经大于30；长度列表已经至少满足3次feed的数量，且为一个feed list的最大允许长度
		// 故可以以此为停止条件，以减少资源的使用
		// 终止条件2：next_time少于stop_time，stop_time设置为了14天前，不断搜索14天前的视频作为feed不符合产品定义，
		// 故作为终止条件
		if len(video_list) > 30 || next_time < stop_time {
			break
		}

		latest_time = next_time
		next_time -= 86400
	}

	return video_list, nil
}

/*
	向后搜索Feed List，后为更接近当前时间的时间点
*/
func SearchFeedLaterInHB(marked_time string, current_time string) (res []VideoInHB, new_marked_time string, err error) {
	marked_time_int, _ := strconv.Atoi(marked_time)
	current_time_int, _ := strconv.Atoi(current_time)

	next_marked_time_int := int64(marked_time_int) + 6*60*60

	video_list := []VideoInHB{}

	for {
		temp, err := FindFeedInHB(fmt.Sprint(marked_time_int), fmt.Sprint(next_marked_time_int))

		if err != nil {
			return video_list, marked_time, err
		}

		video_list = append(video_list, temp...)

		// 终止条件1：视频列表长度已经大于30；长度列表已经至少满足3次feed的数量，且为一个feed list的最大允许长度
		// 故可以以此为停止条件，以减少资源的使用
		// 终止条件2：时间差小于6个小时
		if len(video_list) > 30 || JudgeTimeDiff(next_marked_time_int, fmt.Sprint(current_time_int), 6*60*60) {
			break
		}

		marked_time_int = int(next_marked_time_int)
		next_marked_time_int += 6 * 60 * 60
	}

	return video_list, fmt.Sprint(next_marked_time_int), nil
}
