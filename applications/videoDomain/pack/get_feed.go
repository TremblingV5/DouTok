package pack

import (
	"strconv"

	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageFeedListResp(list []typedef.VideoInHB, errNo *errno.ErrNo, userId int64) (resp *videoDomain.DoutokGetFeedResponse, err error) {
	res := videoDomain.DoutokGetFeedResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
	}

	nextTime := "9999999999"
	var videoList []*entity.Video

	var videoIdList []int64
	validateMap := make(map[int64]bool)
	var userIdList []int64
	validateUserIdMap := make(map[int64]bool)
	for _, v := range list {
		videoId := v.GetId()
		if _, ok := validateMap[videoId]; !ok {
			videoIdList = append(videoIdList, v.GetId())
			validateMap[videoId] = true
		}
		userId := v.GetAuthorId()
		if _, ok := validateUserIdMap[userId]; !ok {
			userIdList = append(userIdList, userId)
			validateUserIdMap[userId] = true
		}
		if v.GetTimestamp() < nextTime {
			nextTime = v.GetTimestamp()
		}
	}

	//userInfoResp, err := rpc.GetUserListByIds(context.Background(), &user.DouyinUserListRequest{
	//	UserList: userIdList,
	//})
	//if err != nil {
	//	return nil, nil
	//}
	//userInfo := userInfoResp.UserList
	//userMap := make(map[int64]*entity.User)
	//for _, v := range userInfo {
	//	userMap[v.Id] = v
	//}

	//isFavoriteResp, err := rpc.IsFavorite(context.Background(), &favorite.DouyinIsFavoriteRequest{
	//	UserId:      userId,
	//	VideoIdList: video_id_list,
	//})
	//if err != nil {
	//	return nil, nil
	//}
	//isFavorite := isFavoriteResp.Result
	//
	//favoriteCountResp, err := rpc.FavoriteCount(context.Background(), &favorite.DouyinFavoriteCountRequest{
	//	VideoIdList: video_id_list,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//favoriteCount := favoriteCountResp.Result

	//commentCountResp, err := rpc.CommentCount(context.Background(), &comment.DouyinCommentCountRequest{
	//	VideoIdList: video_id_list,
	//})
	//if err != nil {
	//	return nil, nil
	//}
	//commentCount := commentCountResp.Result

	for _, v := range list {
		var temp entity.Video

		temp.Id = v.GetId()
		temp.Author = &entity.User{
			Id: v.GetAuthorId(),
		}
		temp.PlayUrl = v.GetVideoUrl()
		temp.CoverUrl = v.GetCoverUrl()
		temp.Title = v.GetTitle()

		//value, ok := favoriteCount[v.GetId()]
		//if ok {
		//	temp.FavoriteCount = value
		//} else {
		//	temp.FavoriteCount = 0
		//}

		//commentCnt, ok := commentCount[v.GetId()]
		//if ok {
		//	temp.CommentCount = commentCnt
		//} else {
		//	temp.CommentCount = 0
		//}

		//isFav, ok := isFavorite[v.GetId()]
		//if ok {
		//	temp.IsFavorite = isFav
		//} else {
		//	temp.IsFavorite = false
		//}

		//temp.Author = userMap[v.GetAuthorId()]

		videoList = append(videoList, &temp)

		if v.GetTimestamp() < nextTime {
			nextTime = v.GetTimestamp()
		}
	}

	res.VideoList = videoList
	nextTimeInt64, _ := strconv.Atoi(nextTime)
	res.NextTime = int64(nextTimeInt64)

	return &res, nil
}
