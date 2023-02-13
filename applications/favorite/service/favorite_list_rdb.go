package service

func QueryFavListInRDB(user_id int64) ([]int64, error) {
	res, err := DoFavorite.Select(
		Favorite.VideoId,
	).Where(
		Favorite.UserId.Eq(user_id),
	).Find()

	if err != nil {
		return nil, err
	}

	result := []int64{}
	for _, v := range res {
		if v.Status == 1 {
			result = append(result, v.VideoId)
		}
	}

	return result, nil
}
