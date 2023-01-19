package main

import (
	"fmt"

	"github.com/TremblingV5/DouTok/applications/publish/service"
)

func InitDb() {
	service.InitDb()
}

func main() {
	InitDb()
	v, _ := service.QueryVideoFromRBDById(1)
	vs, _ := service.QuerySomeVideoFromRDBByIds(1, 22)
	fmt.Println(v)

	for _, v := range vs {
		fmt.Println(v)
	}

	fmt.Println(service.DB)
}
