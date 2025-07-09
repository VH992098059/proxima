package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"proxima/app/einoPackage/source"
	"proxima/app/einoPackage/tools_control"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/schema"
)

func ChatAiModel(ctx context.Context, isNetWork bool, input string) string {
	NewChanMsg := make(chan string, 100000)
	defer close(NewChanMsg)
	var sources []string
	//历史记录
	var fullContent string
	var fullMutilContent []schema.ChatMessagePart
	model, err := source.BuildEinoModel(ctx)
	if err != nil {
		log.Fatal("模型启动失败:", err)
		return ""
	}
	searchTool, err := source.NewTool(ctx)
	if err != nil {
		log.Println("搜索工具初始化失败:", err)
	}
	//启用联网模式
	if isNetWork {
		//使用ddg获取网页信息
		searchReq := &duckduckgo.SearchRequest{
			Query: input,
			Page:  10,
		}
		jsonReq, err := json.Marshal(searchReq)
		if err != nil {
			log.Fatalf("搜索请求序列化失败: %v", err)
		}
		resp, err := searchTool.InvokableRun(ctx, string(jsonReq))
		if err != nil {
			log.Fatal("搜索失败:", err)
		}
		var searchResp duckduckgo.SearchResponse
		if err := json.Unmarshal([]byte(resp), &searchResp); err != nil {
			log.Fatal("解析搜索结果失败:", err)

		}
		for _, result := range searchResp.Results {
			content := tools_control.ExtractMainContent(result.Link)
			if content != "" {
				sources = append(sources, content) //将搜索到的内容添加找到sources
			}
		}
	}
	sources = append(sources, input) //将内容添加
	tools_control.TemplateParams["question"] = sources

	stream, err := model.Stream(ctx, tools_control.TemplateParams)
	if err != nil {
		log.Fatal("输出错误:", err)
	}
	defer stream.Close() // 注意要关闭
	// 处理流式内容
	for {
		chunk, err := stream.Recv()
		if err != nil {
			break
		}

		fmt.Print(chunk.(*schema.Message).Content)
		NewChanMsg <- chunk.(*schema.Message).Content
		fullContent += chunk.(*schema.Message).Content
		fullMutilContent = chunk.(*schema.Message).MultiContent
		_ = chunk.(*schema.Message).Content
	}
	tools_control.TemplateParams["chat_history"] = append(tools_control.TemplateParams["chat_history"].([]*schema.Message), &schema.Message{
		Role:         "assistant",
		Content:      fullContent,
		MultiContent: fullMutilContent,
	})
	return fullContent
}
