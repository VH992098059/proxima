package user_edit_delete

import (
	"context"
	account "proxima/app/user/api/account_edit_delete/v1"

	"proxima/app/gateway/api/user_edit_delete/v1"
)

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	user, err := c.AccountEditDelete.DeleteUser(ctx, &account.DeleteUserReq{Id: int32(req.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserRes{IsDel: int(user.IsDel)}, nil
}
