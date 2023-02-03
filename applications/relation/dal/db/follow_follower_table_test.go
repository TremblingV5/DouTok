package db

import (
	"fmt"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"testing"
)

func TestRelation(t *testing.T) {
	
	//Insert2FollowTable(6, 5)
	//err := Insert2FollowerTable(5, 4)
	//if err != nil {
	//	panic(err)
	//}
	//if err := IsRelation(4, 5); err != nil {
	//	panic(err)
	//}
	//if err := DeleteOnFollowTable(4, 5); err != nil {
	//	panic(err)
	//}
	//if err := DeleteOnFollowerTable(2, 3); err != nil {
	//	panic(err)
	//}
	//err := CancelRelation(100, 200)
	//if err != nil {
	//	panic(err)
	//}
	DB.Where("id>?", 0).Delete(&user.User{})

	DB.Create(&user.User{Id: 1, Name: "test1", FollowCount: 0, FollowerCount: 0})
	DB.Create(&user.User{Id: 2, Name: "test2", FollowCount: 0, FollowerCount: 0})
	DB.Create(&user.User{Id: 3, Name: "test3", FollowCount: 0, FollowerCount: 0})
	DB.Create(&user.User{Id: 4, Name: "test4", FollowCount: 0, FollowerCount: 0})
	DB.Create(&user.User{Id: 5, Name: "test5", FollowCount: 0, FollowerCount: 0})

	if res, err := GetFollowList(5); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
