package user_msg

import (
	"context"
	"fmt"
	"log"
	"proxima/app/user/internal/dao"
	"proxima/app/user/internal/model/entity"
	"proxima/utility/consts"

	"github.com/golang-jwt/jwt/v5"
)

// jwtClaims 定义自定义 JWT 载荷结构
type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

func Info(ctx context.Context, token string) (user *entity.Users, err error) {
	log.Println("Parsing JWT:", token)
	// 解析自定义 Claims
	claims := &jwtClaims{}
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
	// 从 claims.Id 获取用户 ID 并查询
	userID := claims.Id
	err = dao.Users.Ctx(ctx).Where("id", userID).Scan(&user)
	return
}
