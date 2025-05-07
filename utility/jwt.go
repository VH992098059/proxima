package utility

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"proxima/utility/consts"
	"strings"
)

// JwtClaims 定义自定义 JWT 载荷结构
type JwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

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
func GetJWT(ctx context.Context) (token string) {
	// 使用从 Header 绑定的 req.Token，并移除 "Bearer " 前缀
	token = g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	return
}
