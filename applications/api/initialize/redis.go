package initialize

import (
	"fmt"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func InitRedisClient() {

	Client, err := redishandle.InitRedisClient(
		fmt.Sprintf("%s:%d", ViperConfig.Viper.GetString("Redis.Host"), ViperConfig.Viper.GetInt("Redis.Port")),
		ViperConfig.Viper.GetString("Redis.Password"),
		ViperConfig.Viper.GetInt("Redis.Databases.Default"),
	)
	if err != nil {
		panic(err)
	}
	RedisClient = Client
}
