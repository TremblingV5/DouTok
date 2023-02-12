package main

import (
	"github.com/TremblingV5/DouTok/applications/favorite/dal/model"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	var config configStruct.MySQLConfig
	configurator.InitConfig(
		&config, "mysql.yaml",
	)

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
	g.ApplyBasic(model.Favorite{}, model.FavoriteCount{})
	g.ApplyInterface(func() {}, model.Favorite{}, model.FavoriteCount{})

	g.Execute()
}
