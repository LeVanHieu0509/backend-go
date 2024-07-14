package initialize

import (
	"context"
	"fmt"

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
