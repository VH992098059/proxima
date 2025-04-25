// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user_logout

import (
	"context"

	"proxima/app/gateway/api/user_logout/v1"
)

type IUserLogoutV1 interface {
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}
