package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// DB is Redis's client
var DB *redis.Client

// Initialize Redis
func Init(url string, poolSize int) {

	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	opt.DialTimeout = 10 * time.Second
	opt.ReadTimeout = 30 * time.Second
	opt.WriteTimeout = 30 * time.Second
	opt.PoolTimeout = 30 * time.Second

	if poolSize > 0 {
		opt.PoolSize = poolSize
	}

	DB = redis.NewClient(opt)

	_, err = DB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
