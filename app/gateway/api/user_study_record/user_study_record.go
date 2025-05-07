// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user_study_record

import (
	"context"

	"proxima/app/gateway/api/user_study_record/v1"
)

type IUserStudyRecordV1 interface {
	Insert(ctx context.Context, req *v1.InsertReq) (res *v1.InsertRes, err error)
}
