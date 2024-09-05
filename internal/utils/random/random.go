package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitOtp() int {
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + ran.Intn(900000)
	return otp
}
