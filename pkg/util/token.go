package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret = []byte("12345678")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.RegisteredClaims
}
type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.RegisteredClaims
}

// 签发token
func GenerateToken(id uint, userName string, authority int) (token string, err error) {
	//过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	//创建一个claims
	var claims = Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "htt",
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	//映射
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

// 验证token
func ParseToken(token string) (claims *Claims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}

	}

	return nil, err
}

// 签发Emailtoken
func GenerateEmailToken(uId, operationType uint, email, password string) (token string, err error) {
	//过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	//创建一个claims
	var claims = EmailClaims{
		UserID:        uId,
		Email:         email,
		Password:      password,
		OperationType: operationType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "htt",
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	//映射
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

// 验证Emailtoken
func ParseEmailToken(token string) (claims *EmailClaims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}

	}

	return nil, err
}
