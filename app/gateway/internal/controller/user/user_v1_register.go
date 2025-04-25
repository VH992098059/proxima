package user

import (
	"context"
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

	return &v1.RegisterRes{Id: int(id.GetId())}, nil
}
