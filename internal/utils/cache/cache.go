package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	// Lấy giá trị từ Redis dựa trên key
	rs, err := global.Rdb.Get(ctx, key).Result()

	// Xử lý trường hợp Redis trả về lỗi "key không tồn tại"
	if err == redis.Nil {
		return fmt.Errorf(" key %s not found", key)
	} else if err != nil {
		return err
	}

	// Chuyển đổi dữ liệu JSON trong Redis sang đối tượng Go
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("Failed to unmarshal")
	}

	return nil
}
