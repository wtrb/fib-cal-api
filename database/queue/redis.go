package queue

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	rs *redis.Client

	rsChannel string
	ctx       = context.Background()
)

func Add(key, val string) {
	err := rs.Set(ctx, key, val, 0).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

func Get(key string) string {
	val, err := rs.Get(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return val
}

func Publish(key string) {
	err := rs.Publish(ctx, rsChannel, key).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

func init() {
	rsChannel = os.Getenv("REDIS_CHANNEL")

	if rs != nil {
		return
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	addr := fmt.Sprintf("%s:%s", host, port)
	rs = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
