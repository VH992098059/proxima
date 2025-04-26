package account_logout

import (
	"context"
	"log"
	v1 "proxima/app/user/api/account_logout/v1"
	"proxima/app/user/internal/logic/account_logout"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAccountLogoutServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAccountLogoutServer(s.Server, &Controller{})
}

func (*Controller) Logout(ctx context.Context, req *v1.LogoutUserReq) (res *v1.LogoutUserRes, err error) {
	log.Println(req.Token)
	logout, err := account_logout.UserLogout(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &v1.LogoutUserRes{Logout: logout}, nil
}
