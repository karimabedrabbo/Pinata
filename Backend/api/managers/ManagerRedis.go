package managers

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/karimabedrabbo/eyo/api/apputils"
	"log"
)

type Redis struct {
	RedisClient *redis.Client
}

var redisClient *Redis

func SetupRedis() *Redis {
	redisHost, redisPort := apputils.GetRedisHost(), apputils.GetRedisPort()

	//using database 0
	redisUrl := fmt.Sprintf("redis://%s:%s/0", redisHost, redisPort)

	option, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatalf("error connecting to redis instance: %v", err)
	}
	return &Redis{
		RedisClient: redis.NewClient(option),
	}
}

func InitRedis() {
	redisClient = SetupRedis()
}

func GetRedisClient() *Redis {
	return redisClient
}
