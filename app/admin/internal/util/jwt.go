package util

import (
	"github.com/golang-jwt/jwt/v5"
	"mall/ent"
	"time"
)

type Claims struct {
	UserId int32 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *ent.User, d time.Duration) (string, error) {
	// 设置声明（Claims）
	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "admin",
		Subject:   "a",
		ID:        "1",
	}
	claims := Claims{
		UserId:           user.ID,
		RegisteredClaims: registeredClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret")) // TODO: change to real secret
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
