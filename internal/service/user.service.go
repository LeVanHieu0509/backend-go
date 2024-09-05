package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LeVanHieu0509/backend-go/internal/repo"
	"github.com/LeVanHieu0509/backend-go/internal/utils/crypto"
	"github.com/LeVanHieu0509/backend-go/internal/utils/random"
	"github.com/LeVanHieu0509/backend-go/internal/utils/sendto"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
)

/*
- Struct UserService chứa một trường là con trỏ đến UserRepo.
- Hàm khởi tạo NewUserService khởi tạo một đối tượng UserService và gán giá trị cho
userRepo bằng cách tạo một đối tượng UserRepo.
- Phương thức GetInfoUser của UserService gọi phương thức GetInfoUser của UserRepo và trả về kết quả.

*/

// 1. sử dụng struct
// type UserService struct {
// 	//Struct UserService có một trường là userRepo thuộc kiểu con trỏ đến struct UserRepo trong package repo.
// 	userRepo *repo.UserRepo
// }

// // 2. Hàm khởi tạo sử dụng con trỏ
// // Định nghĩa một hàm khởi tạo tên NewUserService trả về con trỏ đến một UserService.
// func NewUserService() *UserService {

// 	// Khởi tạo một đối tượng UserService và gán giá trị cho trường userRepo bằng cách
// 	// gọi hàm NewUserRepo() từ package repo, sau đó trả về con trỏ đến đối tượng này.
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// // Định nghĩa một phương thức tên GetInfoUser cho kiểu UserService
// // Phương thức này nhận một con trỏ us đến UserService và trả về một chuỗi.
// func (us *UserService) GetInfoUser() string {

// 	// Gọi phương thức GetInfoUser từ trường userRepo (là một con trỏ đến UserRepo)
// 	// và trả về kết quả của phương thức này.
// 	return us.userRepo.GetInfoUser()
// }

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	//Struct UserService có một trường là userRepo thuộc kiểu con trỏ đến struct UserRepo trong package repo.
	userRepo repo.IUserRepository
	authRepo repo.IAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, authRepo repo.IAuthRepository) IUserService {
	return &userService{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hashEmail - Không tin ai cả, phải băm cái email này ra và tạo ra 1 kí tự không giống ai đưa vào redis nhắm tăng cường bảo mật
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s", hashEmail)

	// 5. check OTP is available
	// 6. User spam ....

	// 1. Check email exits in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	// 2. Create new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is :::%d\n", otp)
	// 3. Save OTP in redis with expiration time - check thời gian hết hạn
	err := us.authRepo.AddOtp(hashEmail, otp, int64(10*time.Minute))

	if err != nil {
		return response.ErrInvalidOtp
	}
	// 4. Send email OTP
	err = sendto.SendTextEmailOtp([]string{email}, "levanhieu.huflit@gmail.com", strconv.Itoa(otp))
	if err != nil {
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}
