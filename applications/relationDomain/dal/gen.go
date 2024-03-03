package main

import (
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	var config configStruct.MySQLConfig
	if err := configurator.InitConfig(
		&config, "mysql.yaml",
	); err != nil {
		panic(err)
	}

	db, err := mysqlIniter.InitDb(
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	if err != nil {
		panic(err)
	}

	g.UseDB(db)
	g.ApplyBasic(model.Relation{}, model.FollowCount{}, model.FollowerCount{})
	g.ApplyInterface(func() {}, model.Relation{}, model.FollowCount{}, model.FollowerCount{})

	g.Execute()
}
