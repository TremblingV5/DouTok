package configStruct

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const DEFAULT_DATABASE = -1

type Redis struct {
	Host     string `mapstructure:"Host" default:"localhost"`
	Port     string `mapstructure:"Port" default:"6379"`
	Dsn      string `mapstructure:"Dsn" default:"localhost:6379"`
	Password string `mapstructure:"Password" default:"root"`
	// {db name 1}:{db num 1},{db name 2}:{db num 2}
	Databases int `mapstructure:"Databases" default:""`
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

type RedisConfig struct {
	Host      string         `yaml:"host"`
	Port      string         `yaml:"port"`
	Password  string         `yaml:"password"`
	Databases map[string]int `yaml:"databases"`
}
