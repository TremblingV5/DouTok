package service

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/mysqlIniter"
	"github.com/TremblingV5/DouTok/pkg/redisHandle"
)

func InitDb() error {
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
		return err
	}
	DB = db
	return nil
}

func InitRD() error {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)
	redisCaches, err := redishandle.InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)
	if err != nil {
		return err
	}
	RdClient = redisCaches[RdDefault]
	return nil
}
