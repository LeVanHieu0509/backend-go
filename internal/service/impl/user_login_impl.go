package impl

import (
	"context"
	"database/sql"
	"encoding/json"
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
	"github.com/LeVanHieu0509/backend-go/internal/utils/auth"
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
func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	//1. Logic login
	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)

	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	//2. check match pass
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf(" Does not match password")
	}

	//3. check two-factor authentication (nếu kích hoạt rồi thì hệ thống gửi OTP thông qua tài khoản bạn đăn kí, xác nhận OTP qua email thành công thì mới return về token để mà login)
	isTwoFactorEnable, err := s.r.IsTwoFactorEnabled(ctx, uint32(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("")
	}

	if isTwoFactorEnable > 0 {
		//send OTP to in.TwoFactorEmail
		keyUserLoginTwoFactor := crypto.GetHash("2fa:otp:" + strconv.Itoa(int(userBase.UserID)))
		err = global.Rdb.SetEx(ctx, keyUserLoginTwoFactor, "111111", time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()

		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("Set OTP redis failed")
		}

		//send otp via twoFactorEmail
		//get email 2fa
		infoUserTwoFactor, err := s.r.GetTwoFactorMethodByIDAndType(ctx, database.GetTwoFactorMethodByIDAndTypeParams{
			UserID:            uint32(userBase.UserID),
			TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		})

		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("Get two factor method failed")
		}

		log.Println("Send OTP 2FA to Email::", infoUserTwoFactor.TwoFactorEmail)
		go sendto.SendTextEmailOtp([]string{infoUserTwoFactor.TwoFactorEmail.String}, consts.HOST_EMAIL, "111111")
		// go sendto.SenEmail

		out.Message = "Send OTP 2FA to Email"

		return response.ErrCodeSuccess, out, nil

	}

	//4. update status login -> tracking người dùng

	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
		UserAccount:  in.UserAccount,
		UserPassword: in.UserPassword, // Ko cần
	})

	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	log.Println("subToken:", subToken)

	//5. Get user info table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	//6. convert to json
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Convert to json failed")
	}

	//7. Give info user json to redis with key = subToken
	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(consts.TIME_2FA_OTP_REGISTER)*time.Minute).Err()

	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	//8. Create token and refresh token
	out.Token, err = auth.CreateToken(subToken)

	if err != nil {
		return
	}

	return 200, out, nil
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
	//1. hashKey
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	//get OTP
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()

	if err != nil {
		return out, err
	}

	// check OTP user with redis
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

func (s *sUserLogin) UpdatePassword(ctx context.Context, token string, password string) (userId int, err error) {
	// token is already verified
	infoOTP, err := s.r.GetInfoOTP(ctx, token)

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// check is verify OTP

	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeUserOtpNotExists, fmt.Errorf("user OTP not verified")
	}

	// update user base
	userBase := database.AddUserBaseParams{}

	userBase.UserAccount = infoOTP.VerifyKey
	UserSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	userBase.UserSalt = UserSalt
	userBase.UserPassword = crypto.HashPassword(password, UserSalt)

	// Add UserBase to user Base Table
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err := newUserBase.LastInsertId()

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// add user_id to user_info table
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: infoOTP.VerifyKey != ""},
		UserAvatar:           sql.NullString{String: "", Valid: false},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: false},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err = newUserInfo.LastInsertId()

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	return int(user_id), nil
}

// ------------- START TWO FACTOR AUTHENTICATION ------------------- //
func (s *sUserLogin) IsTwoFactorEnable(ctx context.Context, userId int) (codeResult int, rs bool, err error) {

	return 200, true, nil
}
func (s *sUserLogin) SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error) {
	//1. Check isTwoFactorEnable -> true return

	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)

	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}

	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthSetupFailed, fmt.Errorf("Two factor authentication is already enabled")
	}

	//2. create new type authentication
	err = s.r.EnableTwoFactorTypeEmail(ctx, database.EnableTwoFactorTypeEmailParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		TwoFactorEmail:    sql.NullString{String: in.TwoFactorEmail, Valid: true},
	})

	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}

	//3. Send OTP to in.Two Factor Email
	keyUserTwoFactor := crypto.GetHash("2fa:otp:" + strconv.Itoa(int(in.UserId)))

	//4. save otp in cache
	go global.Rdb.Set(ctx, keyUserTwoFactor, "123456", time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()

	return response.ErrCodeSuccess, nil
}
func (s *sUserLogin) VerifyTwoFactorAuth(ctx context.Context, in *model.VerifyTwoFactorAuthInput) (codeResult int, err error) {
	// 1. Check isTwoFactorEnabled
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)

	//
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}

	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("Two-Factor authentication is not enabled")
	}
	// 2. check otp in redis available
	keyUserTwoFactor := crypto.GetHash("2fa:otp:" + strconv.Itoa(int(in.UserId)))

	otpVerifyAuth, err := global.Rdb.Get(ctx, keyUserTwoFactor).Result()
	log.Print("otpVerifyAuth::", otpVerifyAuth)
	if err == redis.Nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	} else if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}

	//3. check OTP
	if otpVerifyAuth != in.TwoFactorCode {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("OTP does not match")
	}

	//4. update status
	err = s.r.UpdateTwoFactorStatus(ctx, database.UpdateTwoFactorStatusParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	//5. remove OTP (khi sử dụng OTP thì nguyên tắc sử dụng OTP là phải xoá)
	_, err = global.Rdb.Del(ctx, keyUserTwoFactor).Result()

	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}

	return 200, nil
}

// ------------- END TWO FACTOR AUTHENTICATION ------------------- //
