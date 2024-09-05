package repo

import (
	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserById(email string) bool
	GetInfoUser() string
}

type UserRepository struct{}

// Implement the interface methods.
func (ur *UserRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??'
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email=?", email).First(&model.GoCrmUser{}).RowsAffected

	return row != NumberNull
}

func (ur *UserRepository) GetUserById(email string) bool {
	panic("unimplemented")
}

func (ur *UserRepository) GetInfoUser() string {
	return "hieu"
}

// NewUserRepository returns a pointer to UserRepository.
func NewUserRepository() IUserRepository {
	return &UserRepository{}
}
