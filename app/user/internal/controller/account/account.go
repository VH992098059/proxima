package account

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "proxima/app/user/api/account/v1"
	"proxima/app/user/api/pbentity"
	"proxima/app/user/internal/logic/account"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAccountServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAccountServer(s.Server, &Controller{})
}

func (*Controller) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {

	id, err := account.Register(ctx, &pbentity.Users{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, gerror.NewCode(gcode.New(10001, err.Error(), nil))
	}
	return &v1.UserRegisterRes{
		Id: int32(id),
	}, nil
}

func (*Controller) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	id, uuid, token, err := account.Login(ctx, req.Username, req.Password)

	if err != nil {
		return nil, gerror.NewCode(gcode.New(10000, err.Error(), nil))
	}
	return &v1.UserLoginRes{
		Id:    id,
		Uuid:  uuid,
		Token: token,
	}, nil
}
