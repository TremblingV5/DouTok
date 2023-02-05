package model

type FavRelation struct {
	UserId  uint64
	VideoId uint64
	Status  bool
}

type VideoCount struct {
	VideoId  uint64
	FavCount uint64
}
