package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义签名用的密钥（服务端保密）
var secretKey = []byte("my_secret_key")

// 1️⃣ 生成 Token
func GenerateToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"uid": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(), // 2小时后过期
		"iss": "my-app",                             // 签发者
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 2️⃣ 解析 Token
func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

// 3️⃣ 演示
func main() {
	tokenStr, _ := GenerateToken(123)
	fmt.Println("生成的Token:", tokenStr)

	claims, _ := ParseToken(tokenStr)
	fmt.Println("解析出的内容:", (*claims)["uid"])
}
