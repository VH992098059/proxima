package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
	"proxima/app/user/internal/controller/account"
	"proxima/app/user/internal/controller/account_edit_delete"
	"proxima/app/user/internal/controller/account_logout"
	"proxima/app/user/internal/controller/user_msg"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "user grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			account.Register(s)
			user_msg.Register(s)
			account_edit_delete.Register(s)
			account_logout.Register(s)
			s.Run()
			return nil
		},
	}
)
