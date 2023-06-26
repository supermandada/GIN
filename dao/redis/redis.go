package redis

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	redis "github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.poolSize"),
	})
	_, err = rdb.Ping(context.Background()).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
