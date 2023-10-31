package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSercrt = []byte("znlxjznlxj")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.RegisteredClaims
}

type EmailClaims struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.RegisteredClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			Issuer:    "LXJ",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSercrt)

	return token, err
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSercrt, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GenerateEmailToken  签发token
func GenerateEmailToken(id uint, email, password string, operationType uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := EmailClaims{
		ID:            id,
		Email:         email,
		Password:      password,
		OperationType: operationType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			Issuer:    "LXJ",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSercrt)

	return token, err
}

// ParseEmailToken 解析 token
func ParseEmailToken(tokenStr string) (*EmailClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSercrt, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*EmailClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
