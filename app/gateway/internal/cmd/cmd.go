package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"proxima/app/gateway/internal/controller/user"
	"proxima/app/gateway/internal/controller/user_edit_delete"
	"proxima/app/gateway/internal/controller/user_logout"
	"proxima/app/gateway/internal/controller/user_msg"
	"proxima/app/gateway/internal/controller/words"
	"proxima/app/gateway/internal/logic/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/gateway", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user.NewV1(),
					words.NewV1(),
				)
				group.Middleware(middleware.Auth)
				group.Bind(
					user_edit_delete.NewV1(),
					user_logout.NewV1(),
				)
				group.Group("/msg", func(group *ghttp.RouterGroup) {
					group.Bind(

						user_msg.NewV1(),
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
