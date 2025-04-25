package user_msg

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	account "proxima/app/user/api/user_msg/v1"
	"strings"

	v1 "proxima/app/gateway/api/user_msg/v1"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// 使用从 Header 绑定的 req.Token，并移除 "Bearer " 前缀
	token := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	log.Println("Forwarding token:", token)
	info, err := c.UserMsgClient.UserInfo(ctx, &account.UserInfoReq{Token: token})
	if err != nil {
		log.Println("gRPC UserInfo error:", err)
		return nil, err
	}
	log.Println("gRPC UserInfo returned:", info)
	return &v1.UserInfoRes{
		Username: info.User.Username,
		Email:    info.User.Email,
		CreateAt: info.User.CreatedAt,
		UpdateAt: info.User.UpdatedAt,
	}, nil
}
