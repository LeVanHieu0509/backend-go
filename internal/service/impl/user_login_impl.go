package impl

import (
	"context"

	"github.com/LeVanHieu0509/backend-go/internal/database"
)

type sUserLogin struct {
	//
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// Implement the IUserLogin interface here
func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePassword(ctx context.Context) error {
	return nil
}
