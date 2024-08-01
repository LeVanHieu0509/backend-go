package repo

// 1. sử dụng struct
type UserRepo struct{}

// GetUserByEmail implements IUserRepository.
func (ur *UserRepo) GetUserByEmail(email string) bool {
	panic("unimplemented")
}

// GetUserById implements IUserRepository.
func (ur *UserRepo) GetUserById(email string) bool {
	panic("unimplemented")
}

// 2. sử dụng con trỏ
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "hieu"
}
