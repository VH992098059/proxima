package model_sse

import (
	"context"
	"fmt"
	"log"
	"proxima/app/einoPackage/api"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

// Client 表示SSE客户端连接
type Client struct {
	Id      string
	Request *ghttp.Request
}

// Service SSE服务
type Service struct {
	clients *gmap.StrAnyMap // 存储所有客户端连接
}

// 创建SSE实例
func New() *Service {
	return &Service{
		clients: gmap.NewStrAnyMap(true),
	}
}

// 创建SSE连接
func (s *Service) Create(r *ghttp.Request) error {
	// 设置SSE响应头
	r.Response.Header().Set("", "chunked")
	r.Response.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")

	clientId := guid.S()
	client := &Client{
		Id:      clientId,
		Request: r,
	}
	s.clients.Set(clientId, client)
	defer s.cleanup(client)

	// 获取用户输入
	message := r.Get("message").String()
	log.Println(message)
	// 启动AI处理并直接处理响应
	return s.processAIResponse(client, message)
}

// 处理AI响应
func (s *Service) processAIResponse(client *Client, message string) error {
	// 发送思考中的提示
	if err := s.sendEvent(client.Request, "thinking", "思考中..."); err != nil {
		return err
	}

	// 直接调用AI模型获取流式响应
	response := api.ChatAiModel(context.Background(), false, message)

	// 处理AI模型的响应
	for _, char := range response {
		select {
		case <-client.Request.Context().Done():
			return nil
		default:
			// 逐字符发送消息到客户端
			if err := s.sendEvent(client.Request, "message", string(char)); err != nil {
				return err
			}
		}
	}

	// 发送完成标记
	return s.sendEvent(client.Request, "done", "会话完成")
}

// 发送事件消息
func (s *Service) sendEvent(r *ghttp.Request, event string, data string) error {
	// 格式化并直接写入SSE消息
	msg := fmt.Sprintf("event: %s\ndata: %s\n\n", event, data)
	r.Response.Write([]byte(msg))

	r.Response.Flush()
	return nil
}

// 清理客户端资源
func (s *Service) cleanup(client *Client) {
	s.clients.Remove(client.Id)
}
