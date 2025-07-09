package user

import (
	"context"
	"log"
	account "proxima/app/user/api/account/v1"

	"proxima/app/gateway/api/user/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	id, err := c.AccountClient.UserRegister(ctx, &account.UserRegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	log.Println("用户注册成功, id:", id.GetId())
	return &v1.RegisterRes{Id: int(id.GetId())}, nil
}
