package user_study_record

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "proxima/app/word/api/user_study_record/v1"
	"proxima/app/word/internal/logic/user_study_record"
)

type Controller struct {
	v1.UnimplementedUserStudyRecordServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserStudyRecordServer(s.Server, &Controller{})
}

// InsertRecord 添加用户学习过的单词
func (*Controller) InsertRecord(ctx context.Context, req *v1.InsertReq) (res *v1.InsertRes, err error) {
	result, err := user_study_record.InsertRecord(ctx, user_study_record.WordProgressSimple{
		WordId:         req.WordId,
		StudyTime:      gtime.Now(),
		MasteryLevel:   req.MasteryLevel,
		ReviewCount:    req.ReviewCount,
		LastReviewTime: gtime.Now(),
		WordLevelId:    req.WordLevelId,
		IsFavorite:     req.IsFavorite,
	}, req.Token)
	if err != nil {
		return nil, err
	}

	return &v1.InsertRes{Msg: result}, nil
}

func (*Controller) UpdateRecord(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SearchRecord(ctx context.Context, req *v1.SearchReq) (res *v1.SearchRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) DeleteRecord(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
