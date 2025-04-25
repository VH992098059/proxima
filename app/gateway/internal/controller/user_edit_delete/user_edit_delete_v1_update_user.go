package user_edit_delete

import (
	"context"
	account "proxima/app/user/api/account_edit_delete/v1"

	"proxima/app/gateway/api/user_edit_delete/v1"
)

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	user, err := c.AccountEditDelete.UpdateUser(ctx, &account.UpdateUserReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserRes{IsOk: int(user.IsOK)}, nil
}
