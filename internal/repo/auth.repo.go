package repo

import (
	"fmt"
	"time"

	"github.com/LeVanHieu0509/backend-go/global"
)

// // 1. sử dụng struct
// type AuthRepo struct{}

// // 2. sử dụng con trỏ
// func NewAuthRepo() *AuthRepo {
// 	return &AuthRepo{}
// }

// func (ur *AuthRepo) FindByUserName() int64 {
// 	var user = po.User{}

// 	db := global.Mdb
// 	result := db.Find(&user, 1)

// 	return result.RowsAffected
// }

//Interface

type IAuthRepository interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type authRepository struct {
}

// GetUserById implements IUserRepository.
func (auth *authRepository) AddOtp(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	fmt.Println(key)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
	// panic("unimplemented")
}

// Tạo instance mới
func NewAuthRepository() IAuthRepository {
	return &authRepository{}
}
