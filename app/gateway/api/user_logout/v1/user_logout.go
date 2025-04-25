package v1

import "github.com/gogf/gf/v2/frame/g"

type LogoutReq struct {
	g.Meta `path:"users/logout" method:"post" sm:"退出登录"`
	Token  string `json:"token"`
}
type LogoutRes struct {
	Logout bool `json:"logout"`
}
