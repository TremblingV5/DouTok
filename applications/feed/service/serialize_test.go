package service

import (
	"fmt"
	"testing"
)

func TestVideoList2String(t *testing.T) {
	video := VideoInHB{
		Id:         1,
		AuthorId:   2,
		AuthorName: "Tom",
	}
	list := []VideoInHB{
		video,
	}
	res := VideoList2String(list)

	fmt.Println(res)
}

func TestString2VideoList(t *testing.T) {
	str := "{\"id\":\"1\",\"author_id\":\"2\",\"author_name\":\"Tom\",\"title\":\"\",\"video_url\":\"\",\"cover_url\":\"\",\"timestamp\":\"\"}"
	list := []string{str}

	res := String2VideoList(list)

	fmt.Println(res)
}
