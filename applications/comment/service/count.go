package service

import (
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func CountComment(video_id ...int64) (map[int64]int64, *errno.ErrNo, error) {
	res, err := DoCommentCnt.Where(
		CommentCnt.Id.In(video_id...),
	).Find()

	if err != nil {
		return nil, &misc.QueryCommentCountErr, err
	}

	count_list := make(map[int64]int64)

	for _, v := range res {
		count_list[v.Id] = v.Number
	}

	return count_list, nil, nil
}

func AddCount(video_id int64) error {
	_, err := DoCommentCnt.Where(
		CommentCnt.Id.Eq(video_id),
	).Update(CommentCnt.Number, CommentCnt.Number.Add(1))

	if err != nil {
		return err
	}

	return nil
}

func ReduceCount(video_id int64) error {
	_, err := DoCommentCnt.Where(
		CommentCnt.Id.Eq(video_id),
	).Update(CommentCnt.Number, CommentCnt.Number.Add(-1))

	if err != nil {
		return err
	}

	return nil
}
