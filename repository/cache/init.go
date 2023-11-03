package cache

import (
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func Redis(RedisAddr, RedisDbName string) {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)

	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	RedisClient = client
}
