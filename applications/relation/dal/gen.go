package main

import (
	"github.com/TremblingV5/DouTok/applications/relation/dal/model"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./applications/relation/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	cfg := dtviper.ConfigInit("", "relation")

	db, err := mysqlIniter.InitDb(
		cfg.Viper.GetString("MySQL.Username"),
		cfg.Viper.GetString("MySQL.Password"),
		cfg.Viper.GetString("MySQL.Host"),
		cfg.Viper.GetString("MySQL.Port"),
		cfg.Viper.GetString("MySQL.Database"),
	)

	if err != nil {
		panic(err)
	}

	g.UseDB(db)
	g.ApplyBasic(model.Relation{}, model.FollowCount{}, model.FollowerCount{})
	g.ApplyInterface(func() {}, model.Relation{}, model.FollowCount{}, model.FollowerCount{})

	g.Execute()
}
