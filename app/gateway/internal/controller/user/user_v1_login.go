package user

import (
	"context"
	"log"
	account "proxima/app/user/api/account/v1"

	"proxima/app/gateway/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, err := c.AccountClient.UserLogin(ctx, &account.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}
	log.Println("用户登录成功, id:", user.Id)
	return &v1.LoginRes{Token: user.Token, Id: user.Id, Uuid: user.Uuid}, nil
}
