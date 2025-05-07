package account_logout

import (
	"context"
	"proxima/utility"
)

// UserLogout 用户退出登录
func UserLogout(ctx context.Context, token string) (logout bool, err error) {
	decryption, err := utility.Decryption(token, &utility.JwtClaims{})
	if err != nil {
		return false, err
	}
	claims := decryption.Claims.(*utility.JwtClaims)
	err = utility.AddBlackTokens(ctx, claims.Username, token)
	if err != nil {
		return false, err
	}
	return true, nil
}
