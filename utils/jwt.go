package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte = []byte("Kninnq1231qwe")

type MyCustomClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

// CheckToken 检查 token 原始串是否合法
func CheckToken(raw string, bd string) (string, error) {
	if bd == "" {
		bd = "Bearer"
	}
	bearer := strings.SplitN(raw, " ", 2)[0]
	token := strings.SplitN(raw, " ", 2)[1]
	if bearer != bd || len(token) < 10 {
		return "", errors.New("token is illegal")
	}
	return token, nil
}

func ParseToken(tokenString string, claims *MyCustomClaims) (resultClaims *MyCustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	resultClaims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		return resultClaims, nil
	}
	return nil, fmt.Errorf("CheckToken: %v", err)
}

func GenToken(uid string) (tokenString string, err error) {
	claims := MyCustomClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
			Issuer:    "niki.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err = token.SignedString(hmacSampleSecret); err != nil {
		return "", err
	}
	return
}
