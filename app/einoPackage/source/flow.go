package source

import (
	"context"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"log"
)

// newLambda component initialization function of node 'Lambda3' in graph 'EinoModel'
func newLambda(ctx context.Context) (lba *compose.Lambda, err error) {
	config := &react.AgentConfig{}
	chatModelIns11, err := newChatModel(ctx)
	if err != nil {
		return nil, err
	}
	config.Model = chatModelIns11
	/*使用工具*/
	toolIns21, err := NewTool(ctx)
	if err != nil {
		return nil, err
	}
	config.ToolsConfig.Tools = []tool.BaseTool{toolIns21}
	ins, err := react.NewAgent(ctx, config)
	if err != nil {
		return nil, err
	}
	log.Println("flow中的NewAgent：", &ins)
	lba, err = compose.AnyLambda(ins.Generate, ins.Stream, nil, nil)
	if err != nil {
		return nil, err
	}
	log.Println("flow中的Lambda", &lba)
	return lba, nil
}
