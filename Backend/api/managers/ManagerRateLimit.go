package managers

import (
	"github.com/karimabedrabbo/eyo/api/constants"
	"github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
)

type RateLimiter struct {
	publicLimiterClient *limiter.Limiter
	identityLimiterClient *limiter.Limiter
}

var rateLimiter *RateLimiter

func SetupRateLimit() *RateLimiter{
	return &RateLimiter{
		publicLimiterClient: generateLimiter("300-D", k.RedisRateLimiterPublicPrefix),
		identityLimiterClient: generateLimiter("10000-D", k.RedisRateLimiterIdentityPrefix),
	}
}

func InitRateLimit() {
	rateLimiter = SetupRateLimit()
}

func GetRateLimit() *RateLimiter {
	return rateLimiter
}

func generateLimiter(format string, prefix string)  *limiter.Limiter {

	rate, err := limiter.NewRateFromFormatted(format)
	if err != nil {
		log.Fatalf("error setting up rate limiter: %v", err)
	}

	store, err := sredis.NewStoreWithOptions(GetRedisClient().RedisClient, limiter.StoreOptions{
		Prefix:   prefix,
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatalf("error setting up redis store on rate limiter: %v", err)
	}

	return limiter.New(store, rate)

}