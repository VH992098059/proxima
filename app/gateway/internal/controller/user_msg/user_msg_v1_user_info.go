package user_msg

import (
	"context"
	"log"
	v1 "proxima/app/gateway/api/user_msg/v1"
	account "proxima/app/user/api/user_msg/v1"
	"proxima/utility"
)

// UserInfo 用户信息获取控制器
func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	token := utility.GetJWT(ctx)
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
