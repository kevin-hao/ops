package utils

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) (tokenString string) {
	expires_seconds := beego.AppConfig.DefaultInt("token::ExpireSeconds", 600)
	KEY := beego.AppConfig.DefaultString("token::Key", "OPS")
	expireAt := time.Now().Add(time.Second * time.Duration(expires_seconds)).Unix()
	mySigningKey := []byte(KEY)

	data := jwt.MapClaims{
		"username": username,
		"exp":      expireAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := token.SignedString(mySigningKey)
	if err == nil {
		return tokenString
	}
	return
}

func ParseToken(tokenString string, Key []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return Key, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil
	}
	return nil, err
}
