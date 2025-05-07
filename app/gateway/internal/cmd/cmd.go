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
	"proxima/app/gateway/internal/controller/user_study_record"
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
				//用户登录注册网关
				group.Bind(
					user.NewV1(),
				)
				//中间件拦截
				group.Middleware(middleware.Auth)
				//用户编辑与注销网关
				group.Bind(
					user_edit_delete.NewV1(),
					user_logout.NewV1(),
				)
				//用户信息网关
				group.Group("/msg", func(group *ghttp.RouterGroup) {
					group.Bind(
						user_msg.NewV1(),
					)
				})
				//学习网关
				group.Group("/learning", func(group *ghttp.RouterGroup) {
					group.Bind(
						words.NewV1(),
						user_study_record.NewV1(),
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
