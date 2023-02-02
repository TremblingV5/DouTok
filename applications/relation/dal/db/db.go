package db

import (
	"fmt"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	mysqlCig := configStruct.MySQLConfig{}
	err := configurator.InitConfig(&mysqlCig, "relation_mysql.yaml")
	if err != nil {
		fmt.Println(err)
	}
	gormDB, err := mysqlIniter.InitDb(mysqlCig.Username, mysqlCig.Password, mysqlCig.Host, mysqlCig.Port, mysqlCig.Database)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = gormDB
	//DB.AutoMigrate(&Follower{}, &Follow{}, &user.User{})
}
func init() {
	Conn()
	DB.AutoMigrate(&Follower{}, &Follow{}, &user.User{})
}
