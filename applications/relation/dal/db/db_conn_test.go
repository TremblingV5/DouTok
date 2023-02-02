package db

import (
	"fmt"
	"testing"
)

func TestDBConn(t *testing.T) {
	Conn()
	fmt.Println(DB.Name())
	//err := AddFollowNum(1)
	//if err != nil {
	//	panic(err)
	//}
	//err = AddFollowerNum(1)
	//if err != nil {
	//	panic(err)
	//}
	if err := DecrFollowNum(1); err != nil {
		panic(err)
	}
	if err := DecrFollowerNum(1); err != nil {
		panic(err)
	}
}
