package context

import (
	"context"
	"errors"
	"log"

	"github.com/LeVanHieu0509/backend-go/internal/utils/cache"
)

type InfoUserUUID struct {
	UserId      int64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	// truy xuất giá trị subjectUUID từ context
	sUUID, ok := ctx.Value("subjectUUID").(string)

	if !ok {
		return "", errors.New("Failed to get subject UUID")
	}

	return sUUID, nil
}

func GetUserIdFromUUID(ctx context.Context) (int64, error) {
	sUUID, err := GetSubjectUUID(ctx)

	if err != nil {
		return 0, err
	}

	var infoUser InfoUserUUID

	// sau khi user gửi request lên thì ssUUID sẽ được truyền tới func này thông qua context
	// từ context sẽ get ra để check trên redis xem còn valid không, nếu ko thì chứng tỏ UUID đã expried
	if err := cache.GetCache(ctx, sUUID, &infoUser); err != nil {
		log.Println("err", err)
		return 0, err
	}

	// Nếu có thì trả về UserID valid
	return infoUser.UserId, nil
}
