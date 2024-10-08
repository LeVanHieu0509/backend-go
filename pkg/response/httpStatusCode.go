package response

const (
	ErrCodeSuccess      = 20001 //Success - Nó là tài liệu nội bộ trong công ty
	ErrCodeParamInvalid = 20003 //Tnvalid
	ErrInvalidToken     = 30001 // token invalid

	ErrCodeUserHasExists = 50001 // User has already register
	ErrInvalidOtp        = 30002
	ErrSendEmailOtp      = 30003

	ErrCodeOtpNotExist = 60009
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "token is invalid",
	ErrCodeUserHasExists: "user has already register",
	ErrInvalidOtp:        "OTP Error",
	ErrSendEmailOtp:      "Fail to send mail OTP",
	ErrCodeOtpNotExist:   "Error Code Otp Not Exist",
}
