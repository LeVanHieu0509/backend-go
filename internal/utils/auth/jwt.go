package auth

import (
	"log"
	"time"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Struct PayloadClaims mở rộng từ jwt.StandardClaims
// là một struct có sẵn trong package jwt
// dùng cho các claim chuẩn trong JWT như exp (thời gian hết hạn), iat (thời gian phát hành), iss (nguồn phát hành), v.v.

type PayloadClaims struct {
	jwt.StandardClaims
}

// Tạo một JWT được ký (signed JWT) với thông tin payload (các claim).
func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}

func CreateToken(uuidToken string) (string, error) {
	timeEx := global.Config.JWT.JWT_EXPIRATION

	if timeEx == "" {
		timeEx = "1h"
	}

	// Chuyển thời gian hết hạn thành time.Duration.
	expiration, err := time.ParseDuration(timeEx)

	if err != nil {
		return "", err
	}

	// Lấy thời gian hiện tại (now) và tính expiresAt dựa trên thời gian hết hạn
	now := time.Now()
	expiresAt := now.Add(expiration)

	/*
		{
		  "exp": 1730259232,
		  "jti": "4ca678ce-1b23-4e1e-bee0-8ad323077d7b",
		  "iat": 1730255632,
		  "iss": "shopdevgo",
		  "sub": "13clitokenb7ebc277-636e-4135-96ad-3e7c15730c94"
		}
	*/

	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(), //U UID ngẫu nhiên để đảm bảo tính duy nhất.
			ExpiresAt: expiresAt.Unix(),    // thời gian hết hạn.
			IssuedAt:  now.Unix(),          // thời gian phát hành token.
			Issuer:    "shopdevgo",         // nguồn phát hành, đặt là "shopdevgo".
			Subject:   uuidToken,           //  chuỗi uuidToken đã được cung cấp.
		},
	})

}

func ParseJwtTokenSubject(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		// để trả về khóa bí mật dùng kiểm tra chữ ký token.
		return []byte(global.Config.JWT.API_SECRET_KEY), nil
	})

	// Kiểm tra xem token đã được phân tích thành công chưa.
	if tokenClaims != nil {
		// Lấy thông tin payload (claims) từ token.
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}

func VerifyTokenSubject(token string) (*jwt.StandardClaims, error) {
	// Phân tích và lấy claims từ token
	claims, err := ParseJwtTokenSubject(token)

	log.Println(err)

	if err != nil {
		return nil, err
	}

	// Token đã hết hạn chưa (exp)?
	if err = claims.Valid(); err != nil {
		return nil, err
	}

	return claims, nil
}
