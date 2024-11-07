package context

import (
	"context"
	"errors"

	"github.com/LeVanHieu0509/backend-go/internal/utils/cache"
)

type InfoUserUUID struct {
	UserId      int64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
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
	if err := cache.GetCache(ctx, sUUID, &infoUser); err != nil {
		return 0, err
	}

	return infoUser.UserId, nil
}
