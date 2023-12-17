package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type JwtCustomClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

var signKey = []byte(viper.GetString("jwt.sign_key"))

func GenerateToken(id uint, name string) (string, error) {
	// 创建一个我们自己的声明
	c := JwtCustomClaims{
		id,
		name, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpires") * time.Minute)), // 过期时间
			Issuer:    "Token",
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(signKey)
}

func ParseToken(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func IsTokenValid(tokenString string) (bool, *JwtCustomClaims) {
	claims, err := ParseToken(tokenString)
	if err != nil || claims.ID == 0 {
		return false, nil
	}
	return true, claims
}
