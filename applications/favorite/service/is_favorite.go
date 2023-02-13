package service

func QueryIsFavorite(user_id int64, video_id_list []int64) (map[int64]bool, error) {
	return nil, nil
}

func QueryIsFavoriteInCache(user_id int64, video_id_list []int64) (res map[int64]bool, is_finished bool, err error) {
	return nil, false, nil
}

func QueryIsFavoriteInRDB(user_id int64, video_id_list []int64, result map[int64]bool) (map[int64]bool, error) {
	return nil, nil
}
