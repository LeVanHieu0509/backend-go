package repo

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

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserById(email string) bool
}

type userRepository struct {
}

// GetUserById implements IUserRepository.
func (u *userRepository) GetUserById(email string) bool {
	panic("unimplemented")
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	panic("unimplemented")
}

// Tạo instance mới
func NewUserRepository() IUserRepository {
	return &userRepository{}
}
