package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserInfoReq struct {
	g.Meta `path:"user/info" method:"get" sm:"获取信息" tags:"用户"`
}

type UserInfoRes struct {
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	CreateAt *timestamppb.Timestamp `json:"create_at"`
	UpdateAt *timestamppb.Timestamp `json:"update_at"`
}
