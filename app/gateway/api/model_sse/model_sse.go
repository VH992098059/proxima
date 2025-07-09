// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model_sse

import (
	"context"

	"proxima/app/gateway/api/model_sse/v1"
)

type IModelSseV1 interface {
	Model(ctx context.Context, req *v1.ModelReq) (res *v1.ModelRes, err error)
}
