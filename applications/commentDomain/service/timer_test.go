package service

import "testing"

func TestUpdateComCountMap(t *testing.T) {
	Init()
	ComCount.Set("1111111111111111112", 22)
	UpdateComCountMap()
}

func TestUpdateComTotalCntMap(t *testing.T) {
	Init()
	ComTotalCount.Set("1111111111111111111", 22)
	UpdateComTotalCntMap()
}
