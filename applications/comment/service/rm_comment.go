package service

import (
	"github.com/TremblingV5/DouTok/applications/comment/dal/model"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func RmComment(user_id int64, comment_id int64) (errno.ErrNo, error) {
	comment, isBelong, err := IsCommentFromUser(user_id, comment_id)
	if err != nil {
		return misc.SystemErr, err
	}
	if !isBelong {
		return misc.CommentNotBelongToUserErr, nil
	}

	rowKey := misc.GetCommentRowKey(comment.VideoId, "0", comment.ConversationId, comment.Timestamp)
	// rowKeyNew := misc.GetCommentRowKey(comment.VideoId, "1", comment.ConversationId, comment.Timestamp)
	// data, err := HBClient.Scan("comment", hbaseHandle.GetFilterByRowKeyPrefix(rowKey)...)

	// if err != nil {
	// 	return misc.HBDataNotFoundErr, err
	// }

	err = HBClient.RmByRowKey("comment", rowKey)
	if err != nil {
		return misc.RmDataFromHBErr, err
	}

	return misc.Success, nil
}

func IsCommentFromUser(user_id int64, comment_id int64) (*model.Comment, bool, error) {
	res, err := DoComment.Where(
		Comment.Id.Eq(comment_id),
	).First()

	if err != nil {
		return nil, false, err
	}

	if res.UserId != user_id {
		return nil, false, nil
	}

	return res, true, nil
}

func RmCommentInRDB(video_id int64) error {
	_, err := DoComment.Where(
		Comment.VideoId.Eq(video_id),
	).Update(Comment.Status, false)

	return err
}
