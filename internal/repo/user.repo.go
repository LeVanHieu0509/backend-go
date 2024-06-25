package repo

// 1. sử dụng struct
type UserRepo struct{}

// 2. sử dụng con trỏ
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "hieu"
}
