package service

import "github.com/TremblingV5/DouTok/applications/user/dal/model"

func QueryUserByIdInRDB(user_id int64) (*model.User, error) {
	user, err := Do.Where(
		User.ID.Eq(uint64(user_id)),
	).First()

	if err != nil {
		return user, err
	}

	return user, nil
}

func QueryUserListByIdInRDB(user_id ...uint64) ([]*model.User, error) {
	user_list, err := Do.Where(
		User.ID.In(user_id...),
	).Find()

	if err != nil {
		return user_list, err
	}

	return user_list, nil
}

func FindUserByUserName(username string) (*model.User, error) {
	res, err := Do.Where(
		User.UserName.Eq(username),
	).First()

	if err != nil {
		return &model.User{}, err
	}

	return res, nil
}
