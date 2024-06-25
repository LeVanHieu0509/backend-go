package service

import "github.com/LeVanHieu0509/backend-go/internal/repo"

/*
- Struct UserService chứa một trường là con trỏ đến UserRepo.
- Hàm khởi tạo NewUserService khởi tạo một đối tượng UserService và gán giá trị cho
userRepo bằng cách tạo một đối tượng UserRepo.
- Phương thức GetInfoUser của UserService gọi phương thức GetInfoUser của UserRepo và trả về kết quả.

*/

// 1. sử dụng struct
type UserService struct {
	//Struct UserService có một trường là userRepo thuộc kiểu con trỏ đến struct UserRepo trong package repo.
	userRepo *repo.UserRepo
}

// 2. Hàm khởi tạo sử dụng con trỏ
// Định nghĩa một hàm khởi tạo tên NewUserService trả về con trỏ đến một UserService.
func NewUserService() *UserService {

	// Khởi tạo một đối tượng UserService và gán giá trị cho trường userRepo bằng cách
	// gọi hàm NewUserRepo() từ package repo, sau đó trả về con trỏ đến đối tượng này.
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// Định nghĩa một phương thức tên GetInfoUser cho kiểu UserService
// Phương thức này nhận một con trỏ us đến UserService và trả về một chuỗi.
func (us *UserService) GetInfoUser() string {

	// Gọi phương thức GetInfoUser từ trường userRepo (là một con trỏ đến UserRepo)
	// và trả về kết quả của phương thức này.
	return us.userRepo.GetInfoUser()
}
