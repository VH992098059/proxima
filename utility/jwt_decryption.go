package utility

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"proxima/utility/consts"
)

func Decryption(token string, claims jwt.Claims) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(consts.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return parsedToken, nil
}
