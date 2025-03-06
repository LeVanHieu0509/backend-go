package initialize

import (
	"context"
	"fmt"
	"log"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/pkg/ultis"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// context để kiểm soát go routing
var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database, // use default DB
		PoolSize: r.PoolSize, // Số lượng kết nối tối ra( có 10 connect cho 10 cpu sử dụng)
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}
	global.Logger.Info("Initializing Redis Successfully!")
	fmt.Println("Init Redis is running")
	global.Rdb = rdb

	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()

	ultis.HandleErr(err, "Error redis setting: ")

	value, err := global.Rdb.Get(ctx, "score").Result()

	ultis.HandleErr(err, "Error redis setting")

	global.Logger.Info("value score is: ", zap.String("score", value))

}

func InitRedisSentinel() {
	r := global.Config.Redis
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",                                                        // The master name managed by Sentinel
		SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"}, // List of Sentinel addresses
		DB:            r.Database,                                                        // Default database
		Password:      r.Password,                                                        // Password, if Redis has one
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

	fmt.Println("Got value from Redis:", val)
}

func InitRedisSentinelViblo() {
	fmt.Println("-----start connec redis-------")
	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"127.0.0.1:26376", "127.0.0.1:26375", "127.0.0.1:26379"},
		DB:            0, // use default DB
		Password:      "CkkUdLPMT",
	})

	// Check the connection
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	fmt.Println("Got value from Redis:", val)
}
