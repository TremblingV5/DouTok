package main

import (
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/bytedance/gopkg/util/logger"
	"gorm.io/gen"
)

type Method interface {
	//select * from @@table
	GetAllData() []gen.T
}
type Querier interface {
	//select * from @@table where user_id = @user_id
	QueryWihtUserId(user_id int64) (gen.T, error)
	//update follow_follower_count set follow_count = follow_count+1 where user_id = @id
	AddFollowCount(id int64) error
	//update follow_follower_count set follower_count = follower_count+1 where user_id = @id
	AddFollowerCount(id int64) error
	//update follow_follower_count set follow_count = follow_count-1 where user_id = @id
	DecrFollowCount(id int64) error
	//update follow_follower_count set follower_count = follower_count-1 where user_id = @id
	DecrFollowerCount(id int64) error
}

func main() {
	//读取配置
	v, err := conf.InitConfig("./config", "relation")
	if err != nil {
		logger.Fatal(err)
	}
	//连接数据库
	if err := db.Conn(v); err != nil {
		logger.Fatal(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./applications/relation/dal/query",

		Mode: gen.WithoutContext | gen.WithQueryInterface | gen.WithDefaultQuery,
	})
	g.UseDB(db.DB)
	m := g.GenerateModel("follow_follower_count")
	g.ApplyInterface(func(Querier) {}, m)
	//g.ApplyBasic(m)

	g.WithDataTypeMap(map[string]func(string) string{
		"int": func(s string) string {
			return "int64"
		},
	})
	g.Execute()
	//query.SetDefault(db.DB)
	//query.FollowFollowerCount.AddFollowCount(1)
	//query.FollowFollowerCount.AddFollowerCount(1)
	//query.User.MyTest()

	//r := query.User.GetById(1)
	//
	//fmt.Println(r.Name)
	//fmt.Println(c)

}
