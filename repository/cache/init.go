package cache

import (
	"MyMall/config"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func Redis() {
	db, _ := strconv.ParseUint(config.RedisDbName, 10, 64)

	client := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
		DB:   int(db),
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	RedisClient = client
}
