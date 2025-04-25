package user_msg

import (
	"context"
	"proxima/app/user/api/pbentity"
	v1 "proxima/app/user/api/user_msg/v1"
	"proxima/app/user/internal/logic/user_msg"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAccountMsgServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAccountMsgServer(s.Server, &Controller{})
}

func (*Controller) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	token := req.GetToken()
	user, err := user_msg.Info(ctx, token)
	if err != nil {
		return nil, err
	}

	return &v1.UserInfoRes{
		User: &pbentity.Users{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt.Time),
			UpdatedAt: timestamppb.New(user.UpdatedAt.Time),
		},
	}, nil
}
