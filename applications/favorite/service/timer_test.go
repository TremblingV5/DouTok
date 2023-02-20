package service

import "testing"

func TestUpdateFavMap(t *testing.T) {
	Init()
	FavCount.Set("111111111111111111", 22)
	UpdateFavMap()
}

func TestUpdateFavCntMap(t *testing.T) {
	Init()
	FavTotalCount.Set("111111111111111111", 22)
	UpdateFavMap()
}
