package account_logout

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"proxima/utility"
)

// jwtClaims 定义自定义 JWT 载荷结构
type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

// UserLogout 用户退出登录
func UserLogout(ctx context.Context, token string) (logout bool, err error) {
	decryption, err := utility.Decryption(token, &jwtClaims{})
	if err != nil {
		return false, err
	}
	claims := decryption.Claims.(*jwtClaims)
	err = utility.AddBlackTokens(ctx, claims.Username, token)
	if err != nil {
		return false, err
	}
	return true, nil
}
