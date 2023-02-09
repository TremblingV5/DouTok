package redishandle

import (
	"context"
	"testing"
	"time"

	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
)

func TestSet(t *testing.T) {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)

	redisCaches, err := InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)

	if err != nil {
		panic(err)
	}

	if err := redisCaches[constants.DbDefault].Set(context.Background(), "test_key_0001", "test_value_0001", 60*time.Second); err != nil {
		panic(err)
	}
}

type TestOjb struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestRedisClient_SetObj(t *testing.T) {
	var config configStruct.RedisConfig
	configurator.InitConfig(
		&config, "redis.yaml",
	)

	redisCaches, err := InitRedis(
		config.Host+":"+config.Port, config.Password, config.Databases,
	)
	if err != nil {
		panic(err)
	}

	stu := TestOjb{Name: "xiaoming", Age: 15}
	redisCaches[constants.DbDefault].SetObj(context.Background(), "user2", &stu, 5*time.Minute)
	stu2 := TestOjb{}
	redisCaches[constants.DbDefault].GetObj(context.Background(), "user2", &stu2)
	println(stu2)

	//mp, _ := misc.Struct2Map(&stu)
	//println("name = ", mp["name"].(string))
	//println("age = ", mp["age"].(float64))
	//redisCaches[constants.DbDefault].Client.HSet(context.Background(), "user", mp)
	//name := redisCaches[constants.DbDefault].Client.HGet(context.Background(), "user", "name").String()
	//age := redisCaches[constants.DbDefault].Client.HGet(context.Background(), "user", "age").String()
	//println(name, age)
}
