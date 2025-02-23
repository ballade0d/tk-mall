package util

import (
	"github.com/golang-jwt/jwt/v5"
	"mall/ent"
	"time"
)

type Claims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *ent.User, d time.Duration) (string, error) {
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

func VerifyJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // TODO: change to real secret
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
