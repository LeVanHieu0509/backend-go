package initialize

import (
	"context"
	"fmt"
	"log"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port), // 55000
		Password: r.Password,                           // no password set
		DB:       r.Database,                           // use default DB
		PoolSize: 10,                                   //
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}

	// fmt.Println("Initializing Redis Successfully")
	global.Logger.Info("Initializing Redis Successfully")
	global.Rdb = rdb
	// redisExample()
}

// advanced
func InitRedisSentinel() {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster", // Tên master do Sentinel quản lý
		SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},
		DB:            0,        // Sử dụng database mặc định
		Password:      "123456", // Nếu Redis có mật khẩu, điền vào đây
	})

	// Check the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis Sentinel: %v", err)
	}

	fmt.Println("Connected to Redis Sentinel successfully!")

	// Try setting and getting a value
	err = rdb.Set(ctx, "test_key", "Hello Redis Sentinel!", 0).Err()
	if err != nil {
		log.Fatalf("Error setting key: %v", err)
	}

	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		log.Fatalf("Error getting key: %v", err)
	}

	fmt.Println("Value of test_key:", val)

	global.Logger.Info("Initializing RedisSentinel Successfully")
	global.Rdb = rdb
	// redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err))
		return
	}

	global.Logger.Info("value score is::", zap.String("score", value))
}
