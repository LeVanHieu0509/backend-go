package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	rateLimit := &RateLimiter{
		globalRateLimiter:         rateLimiter("100-S"),
		publicAPIRateLimiter:      rateLimiter("80-S"),
		userPrivateAPIRateLimiter: rateLimiter("50-S"),
	}

	return rateLimit
}

func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(global.Rdb, limiter.StoreOptions{
		Prefix:          "rate-limiter", //u:1001
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})

	if err != nil {
		return nil
	}

	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}
	instance := limiter.New(store, rate)

	return instance
}

// Global API Limiter
func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := "global" //unit
		log.Println("Global ---> ")

		limitContext, err := rl.globalRateLimiter.Get(ctx, key)

		if err != nil {
			fmt.Println("Failed to check rate limit GLOBAL", err)
			ctx.Next()
			return
		}

		if limitContext.Reached {
			log.Println("Rate Limit breached Global %s", key)
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached GLOBAL, try latter"})
			return
		}
	}
}

// Public API Limiter
func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlPath := ctx.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath != nil {
			log.Println("Client IP ---> ", ctx.ClientIP())
			key := fmt.Sprintf("%s-%s", "111-222-333-44", urlPath)

			limitContext, err := rateLimitPath.Get(ctx, key)

			if err != nil {
				fmt.Println("Failed to check rate limit GLOBAL", err)
				ctx.Next()
				return
			}
			if limitContext.Reached {
				log.Println("Rate Limit breached %s", key)
				ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try latter"})
				return
			}
		}
		ctx.Next()
	}
}

// Private API Limiter
func (rl *RateLimiter) UserAndPrivateAPIRateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlPath := ctx.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath != nil {
			userId := 1001

			key := fmt.Sprintf("%d-%s", userId, urlPath)

			limitContext, err := rateLimitPath.Get(ctx, key)

			if err != nil {
				fmt.Println("Failed to check rate limit GLOBAL", err)
				ctx.Next()
				return
			}
			if limitContext.Reached {
				log.Println("Rate Limit breached %s", key)
				ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try latter"})
				return
			}
		}
		ctx.Next()
	}
}

func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	log.Println("urlPath", urlPath)
	if urlPath == "v1/2024/user/login" || urlPath == "/ping/80" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "v1/2024/user/info" || urlPath == "/ping/50" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}
