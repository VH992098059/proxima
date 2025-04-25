package account_edit_delete

import (
	"context"
	v1 "proxima/app/user/api/account_edit_delete/v1"
	"proxima/app/user/api/pbentity"
	"proxima/app/user/internal/logic/account_edit_delete"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAccountEditDeleteServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAccountEditDeleteServer(s.Server, &Controller{})
}

func (*Controller) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	isOk, err := account_edit_delete.UpdateUser(ctx, &pbentity.Users{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserRes{IsOK: int32(isOk)}, nil
}

func (*Controller) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	isDel, err := account_edit_delete.DeleteUser(ctx, &pbentity.Users{
		Id: uint32(req.Id),
	})
	if err != nil {
		return &v1.DeleteUserRes{IsDel: int32(isDel)}, err
	}
	return &v1.DeleteUserRes{IsDel: int32(isDel)}, nil
}
