package service

import (
	"encoding/json"
	"fmt"

	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
)

func VideoList2String(list []typedef.VideoInHB) []string {
	res := []string{}

	for _, v := range list {
		r, err := json.Marshal(v)
		if err != nil {
			continue
		}
		res = append(res, string(r))
	}

	return res
}

func String2VideoList(list []string) []typedef.VideoInHB {
	res := []typedef.VideoInHB{}

	for _, v := range list {
		temp := typedef.VideoInHB{}
		err := json.Unmarshal([]byte(v), &temp)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, temp)
	}

	return res
}

func HBaseType2RPCType(v typedef.VideoInHB) *entity.Video {
	return &entity.Video{
		Id: v.GetId(),
		Author: &entity.User{
			Id: v.GetAuthorId(),
		},
		Title:    v.GetTitle(),
		PlayUrl:  v.GetVideoUrl(),
		CoverUrl: v.GetCoverUrl(),
	}
}
