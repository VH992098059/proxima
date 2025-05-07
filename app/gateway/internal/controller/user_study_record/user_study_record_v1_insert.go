package user_study_record

import (
	"context"
	"proxima/app/gateway/api/user_study_record/v1"
	v2 "proxima/app/word/api/user_study_record/v1"
	"proxima/utility"
)

func (c *ControllerV1) Insert(ctx context.Context, req *v1.InsertReq) (res *v1.InsertRes, err error) {
	token := utility.GetJWT(ctx)
	_, err = c.UserStudyRecord.InsertRecord(ctx, &v2.InsertReq{
		WordId:       req.WordId,
		MasteryLevel: req.MasteryLevel,
		ReviewCount:  req.ReviewCount,
		WordLevelId:  req.WordLevelId,
		IsFavorite:   req.IsFavorite,
		Token:        token,
	})
	if err != nil {
		return nil, err
	}

	return &v1.InsertRes{}, nil
}
