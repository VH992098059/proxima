package model_sse

import (
	"context"
	v1 "proxima/app/gateway/api/model_sse/v1"
	"proxima/app/gateway/internal/logic/model_sse"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct {
	service *model_sse.Service
}

func NewV1() *Controller {
	return &Controller{
		service: model_sse.New(),
	}
}

func (c *Controller) Chat(ctx context.Context, req *v1.ModelReq) (res *v1.ModelRes, err error) {
	err = c.service.Create(g.RequestFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	return &v1.ModelRes{}, nil
}
