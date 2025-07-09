package source

import (
	"context"
	"log"

	"github.com/cloudwego/eino/compose"
)

func BuildEinoModel(ctx context.Context) (r compose.Runnable[any, any], err error) {
	const (
		CustomChatTemplate2 = "CustomChatTemplate2"
		Lambda3             = "Lambda3"
	)
	g := compose.NewGraph[any, any]()
	customChatTemplate2KeyOfChatTemplate, err := newChatTemplate(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddChatTemplateNode(CustomChatTemplate2, customChatTemplate2KeyOfChatTemplate)
	lambda3KeyOfLambda, err := newLambda(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddLambdaNode(Lambda3, lambda3KeyOfLambda)
	_ = g.AddEdge(compose.START, CustomChatTemplate2)
	_ = g.AddEdge(Lambda3, compose.END)
	_ = g.AddEdge(CustomChatTemplate2, Lambda3)
	r, err = g.Compile(ctx, compose.WithGraphName("EinoModel"))
	log.Println("orchestration：", r)
	if err != nil {
		return nil, err
	}
	return r, err
}
