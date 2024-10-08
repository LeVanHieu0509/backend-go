package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/consts"
	"github.com/LeVanHieu0509/backend-go/internal/database"
	"github.com/LeVanHieu0509/backend-go/internal/model"
	"github.com/LeVanHieu0509/backend-go/internal/utils"
	"github.com/LeVanHieu0509/backend-go/internal/utils/crypto"
	"github.com/LeVanHieu0509/backend-go/internal/utils/random"
	"github.com/LeVanHieu0509/backend-go/internal/utils/sendto"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
	"github.com/redis/go-redis/v9"
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

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	//1. Hash Email
	fmt.Printf("VerifyKey: %v\n", in.VerifyKey)
	fmt.Printf("VerifyType: %d\n", in.VerifyType)
	fmt.Printf("VerifyPurpose: %v\n", in.VerifyPurpose)

	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	fmt.Printf("hashKey: %v\n", hashKey)

	//2. Check user exist in user Base
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)

	if err != nil {
		return response.ErrCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("User has already registered")
	}

	//3. Create OTP
	userKey := utils.GetUserKey(hashKey)

	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist!")
	case err != nil:
		fmt.Println("Get failed::", err)
	case otpFound != "":
		return response.ErrCodeOtpNotExist, fmt.Errorf("")
	}

	// 4. Generate New OTP
	otpNew := random.GenerateSixDigitOtp()

	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("otpNew: %v\n", otpNew)

	//5. SAVE OTP IN REDIS
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(time.Minute*time.Duration(consts.TIME_OTP_REGISTER))).Err()

	if err != nil {
		return response.ErrInvalidOtp, err
	}

	//6. Send OTP
	switch in.VerifyType {
	case consts.EMAIL:
		err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))

		if err != nil {
			return response.ErrSendEmailOtp, err
		}

		//7. save OTP to DB
		result, err := s.r.InsertOTPVerify(ctx,
			database.InsertOTPVerifyParams{
				VerifyIDOtp:   strconv.Itoa(otpNew),
				VerifyKey:     in.VerifyKey,
				VerifyKeyHash: hashKey,
				VerifyType:    sql.NullInt32{Int32: 1, Valid: true}})

		if err != nil {
			return response.ErrSendEmailOtp, err
		}

		//8 . get Last Id
		lastIdVerifyUser, err := result.LastInsertId()

		if err != nil {
			return response.ErrSendEmailOtp, err
		}

		log.Println("lastIdVerifyUser", lastIdVerifyUser)

		return response.ErrCodeSuccess, nil

	}

	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	//get OTP
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()

	if err != nil {
		return out, err
	}

	if in.VerifyCode != otpFound {
		// if wrong 3t /minute

		return out, fmt.Errorf("OTP not match")
	}

	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)

	if err != nil {
		return out, err
	}

	//update status verify
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)

	if err != nil {
		return out, err

	}

	//out put
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"

	return out, err
}

func (s *sUserLogin) UpdatePassword(ctx context.Context) error {
	return nil
}
