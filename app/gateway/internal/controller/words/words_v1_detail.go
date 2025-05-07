package words

import (
	"context"
	v2 "proxima/app/word/api/words/v1"

	"proxima/app/gateway/api/words/v1"
)

// Detail 单词获取控制器
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	get, err := c.WordClient.Get(ctx, &v2.GetReq{
		Start: uint32(req.Start) - 1,
		Limit: uint32(req.Limit),
	})
	if err != nil {
		return nil, err
	}
	return &v1.DetailRes{
		WordList: get.Words,
	}, nil
}
