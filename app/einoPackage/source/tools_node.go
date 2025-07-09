package source

import (
	"context"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/ddgsearch"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
)

func NewTool(ctx context.Context) (bt tool.InvokableTool, err error) {
	// TODO Modify component configuration here.
	config := &duckduckgo.Config{
		MaxResults: 3,
		Region:     ddgsearch.RegionWT,
		DDGConfig: &ddgsearch.Config{
			Cache:      true,
			MaxRetries: 4,
			Proxy:      "http://127.0.0.1:10808",
		},
	}
	bt, err = duckduckgo.NewTool(ctx, config)
	if err != nil {
		return nil, err
	}
	return bt, nil
}
