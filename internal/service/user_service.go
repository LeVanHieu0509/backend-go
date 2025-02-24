package service

import (
	"context"

	"github.com/LeVanHieu0509/backend-go/internal/model"
)

// Các phương thức này đều nhận context.Context (biến ctx), giúp kiểm soát luồng xử lý, truyền dữ liệu và timeout giữa các luồng

// Việc sử dụng các interface giúp tách biệt giữa logic nghiệp vụ và logic triển khai,
// đồng thời tạo sự linh hoạt trong việc thay đổi cách triển khai mà không ảnh hưởng đến phần còn lại của ứng dụng.
type (
	//...interface
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error)
		UpdatePassword(ctx context.Context, token string, password string) (userId int, err error)

		//two-factor authentication
		IsTwoFactorEnable(ctx context.Context, userId int) (codeResult int, rs bool, err error)
		SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error)
		VerifyTwoFactorAuth(ctx context.Context, in *model.VerifyTwoFactorAuthInput) (codeResult int, err error)
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserInfo  IUserInfo
	localUserLogin IUserLogin
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("Implement localUserAdmin not found for interface IUserAdmin")
	}

	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("Implement localUserInfo not found for interface IUserInfo")
	}

	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("Implement localUserLogin not found for interface IUserLogin")
	}

	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
