package user_logout

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	account "proxima/app/user/api/account_logout/v1"
	"strings"

	"proxima/app/gateway/api/user_logout/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	// 使用从 Header 绑定的 req.Token，并移除 "Bearer " 前缀
	token := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	log.Println("需要退出登录token:", token)
	logout, err := c.AccountLogout.Logout(ctx, &account.LogoutUserReq{Token: token})
	if err != nil {
		return nil, err
	}
	return &v1.LogoutRes{Logout: logout.Logout}, nil
}
