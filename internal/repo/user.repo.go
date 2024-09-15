package repo

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserById(email string) bool
	GetInfoUser() string
}

type UserRepository struct {
	sqlc *database.Queries
}

// Implement the interface methods.
func (ur *UserRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??'

	// Thay vì viết như này bằng goose thì sẽ thay thế bằng cách mysqlc
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email=?", email).First(&model.GoCrmUser{}).RowsAffected

	// return row != NumberNull

	//cách 2 sử dụng mysqlc
	user, err := ur.sqlc.GetUserByEmailSQLC(ctx, email)
	fmt.Printf("UsrID::%v", user)

	if err != nil {
		return false
	}

	return user.UsrID != 0

}

func (ur *UserRepository) GetUserById(email string) bool {
	panic("unimplemented")
}

func (ur *UserRepository) GetInfoUser() string {
	return "hieu"
}

// NewUserRepository returns a pointer to UserRepository.
func NewUserRepository() IUserRepository {
	return &UserRepository{
		sqlc: database.New((global.Mdbc)),
	}
}
