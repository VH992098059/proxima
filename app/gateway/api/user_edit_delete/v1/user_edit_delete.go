package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateUserReq struct {
	g.Meta   `path:"/user/update" method:"put" sm:"修改信息" tags:"用户"`
	Username string `json:"username" v:"required"`
	Password string `json:"password" v:"required"`
	Email    string `json:"email" v:"required"`
}
type UpdateUserRes struct {
	IsOk int `json:"isOK"`
}
type DeleteUserReq struct {
	g.Meta `path:"/user/delete" method:"delete" sm:"删除用户" tags:"用户"`
	Id     int `json:"id" v:"required"`
}
type DeleteUserRes struct {
	IsDel int `json:"isDel" `
}
