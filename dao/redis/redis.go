package redis

import (
	"context"
	"fmt"
	"gin-project/settings"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func Init(conf *settings.RedisConfig) (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
		PoolSize: conf.PoolSize,
	})

	_, err = rdb.Ping(ctx).Result()

	return
}

func Close() {
	rdb.Close()
}
