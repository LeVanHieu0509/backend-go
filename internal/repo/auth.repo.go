package repo

import (
	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/po"
)

// 1. sử dụng struct
type AuthRepo struct{}

// 2. sử dụng con trỏ
func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

func (ur *AuthRepo) FindByUserName() int64 {
	var user = po.User{}

	db := global.Mdb
	result := db.Find(&user, 1)

	return result.RowsAffected
}
