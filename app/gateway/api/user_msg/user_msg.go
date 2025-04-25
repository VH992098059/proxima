// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user_msg

import (
	"context"

	"proxima/app/gateway/api/user_msg/v1"
)

type IUserMsgV1 interface {
	UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error)
}
