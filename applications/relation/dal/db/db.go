package db

import (
	"github.com/TremblingV5/DouTok/applications/relation/dal/query"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn(v *viper.Viper) error {
	host := v.GetString("mysql.host")
	port := v.GetString("mysql.port")
	password := v.GetString("mysql.password")
	username := v.GetString("mysql.username")
	database := v.GetString("mysql.database")

	gormDB, err := mysqlIniter.InitDb(username, password, host, port, database)
	if err != nil {
		return err
	}
	DB = gormDB
	if err := DB.AutoMigrate(&Follower{}, &Follow{}, &user.User{}); err != nil {
		return err
	}
	query.SetDefault(DB)

	return nil

}
