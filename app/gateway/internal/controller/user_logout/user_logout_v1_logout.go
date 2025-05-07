package user_logout

import (
	"context"
	"log"
	"proxima/app/gateway/api/user_logout/v1"
	account "proxima/app/user/api/account_logout/v1"
	"proxima/utility"
)

// Logout 用户退出登录控制器
func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	token := utility.GetJWT(ctx)
	log.Println("需要退出登录token:", token)
	logout, err := c.AccountLogout.Logout(ctx, &account.LogoutUserReq{Token: token})
	if err != nil {
		return nil, err
	}
	return &v1.LogoutRes{Logout: logout.Logout}, nil
}
