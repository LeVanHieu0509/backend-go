package service

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/internal/repo"
	ty "github.com/LeVanHieu0509/backend-go/types_custom"
)

type AuthService struct {
	authRepo *repo.AuthRepo
}

func NewAuthService() *AuthService {
	return &AuthService{
		authRepo: repo.NewAuthRepo(),
	}
}

func (auth *AuthService) Login(data ty.LoginReq) string {
	/*
		step:
		1. Mỗi User khi login đều phải có 2 key: publicKey và privateKey
		2. privateKey là key secret của AccessToken
		3. Kiểm tra user name, password
		4. Lấy cặp key rsa từ database của user đó để gen ra token trả về cho client kèm thông tin

	*/
	foundUser := auth.authRepo.FindByUserName()

	if foundUser == 0 {
		fmt.Println("User not found")
		return "0"
	}

	fmt.Println("User found:", foundUser)
	return "1"
}

func (auth *AuthService) SignUp(data ty.SignUpRequest) string {
	/*
		1. Check User
		2. Hash password
		3. Save info database
		3. Gen 2 key: publicKey và privateKey
		4. Create Keystore để lưu 2 key xuống database
		4. Gen token dựa vào privateKey để trả về cho user
	*/

	return "1"
}

func (auth *AuthService) Logout(data ty.SignUpRequest) string {
	/*
		1. Check header, get key store
		2. Check resfresh token
		3. Xoá key in redis
	*/
	return "1"
}

func (auth *AuthService) RefreshToken(data ty.SignUpRequest) string {
	/*
		1. Check header
		2. check key store
		3. verify refresh token
		4. create token
	*/
	return "1"
}
