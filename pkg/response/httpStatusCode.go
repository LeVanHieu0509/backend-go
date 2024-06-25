package response

const (
	ErrCodeSuccess      = 20001 //Success - Nó là tài liệu nội bộ trong công ty
	ErrCodeParamInvalid = 20003 //Tnvalid
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
}
