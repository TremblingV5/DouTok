package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/query"
	"github.com/TremblingV5/DouTok/pkg/constants"
)

func ReadFollowCountFromDB(user_id int64) (error, int64) {
	res, err := query.FollowCount.Where(query.FollowCount.UserId.Eq(user_id)).First()
	if err != nil {
		return err, 0
	}
	return nil, res.Number
}

func UpdateFollowCountFromDB(user_id int64, op int64) error {
	res, err := query.FollowCount.Where(
		query.FollowCount.UserId.Eq(user_id),
	).Find()
	if err != nil {
		return err
	}
	if len(res) > 0 {
		// 已经存在
		_, err := query.FollowCount.Where(
			query.FollowCount.UserId.Eq(user_id),
		).Update(query.FollowCount.Number, query.FollowCount.Number.Add(op))
		return err
	} else {
		err := query.FollowCount.Create(
			&model.FollowCount{
				UserId: user_id,
				Number: op,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateFollowerCountFromDB(user_id int64, op int64) error {
	res, err := query.FollowerCount.Where(
		query.FollowerCount.UserId.Eq(user_id)).Find()
	if err != nil {
		return err
	}
	if len(res) > 0 {
		_, err := query.FollowerCount.Where(
			query.FollowerCount.UserId.Eq(user_id),
		).Update(query.FollowerCount.Number, query.FollowerCount.Number.Add(op))
		return err
	} else {
		err := query.FollowerCount.Create(
			&model.FollowerCount{
				UserId: user_id,
				Number: op,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteFollowCountCache(user_id string) error {
	_, err := RedisClient.HDel(context.Background(), user_id, constants.FollowCount).Result()
	if err != nil {
		return err
	}
	return nil
}

func ReadFollowerCountFromDB(user_id int64) (error, int64) {
	res, err := query.FollowerCount.Where(query.FollowerCount.UserId.Eq(user_id)).First()
	if err != nil {
		return err, 0
	}
	return nil, res.Number
}

func DeleteFollowerCountCache(user_id string) error {
	_, err := RedisClient.HDel(context.Background(), user_id, constants.FollowCount).Result()
	if err != nil {
		return err
	}
	return nil
}
