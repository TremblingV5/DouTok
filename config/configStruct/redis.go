package configStruct

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const DEFAULT_DATABASE = -1

type Redis struct {
	Host string `env:"REDIS_HOST" envDefault:"localhost" configPath:"Redis.Host"`
	Port string `env:"REDIS_PORT" envDefault:"6379" configPath:"Redis.Port"`
	//Dsn      string `env:"REDIS_DSN" envDefault:"localhost:6379" configPath:"Redis.Dest"`
	Password string `env:"REDIS_PASSWORD" envDefault:"root" configPath:"Redis.Password"`
	// {db name 1}:{db num 1},{db name 2}:{db num 2}
	Databases int `mapstructure:"Databases" default:"0"`
}

func (r *Redis) InitRedisClient(database int) *redis.Client {
	if database == DEFAULT_DATABASE {
		database = r.Databases
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password: r.Password,
		DB:       database,
		PoolSize: 20,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return client
}
