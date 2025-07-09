package source

import (
	"context"
	"fmt"
	"log"
	"proxima/app/einoPackage/tools_control"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type ChatTemplateImpl struct {
	config *ChatTemplateConfig
}

type ChatTemplateConfig struct {
	Role       schema.RoleType
	System     schema.RoleType
	FormatType schema.FormatType
	Templates  []schema.MessagesTemplate
}

// newChatTemplate component initialization function of node 'CustomChatTemplate2' in graph 'EinoModel'
func newChatTemplate(ctx context.Context) (ctp prompt.ChatTemplate, err error) {
	config := &ChatTemplateConfig{
		/*使用FString，支持Python风格字符串格式化*/
		Role:       schema.User,
		System:     schema.System,
		FormatType: schema.FString,
		Templates: []schema.MessagesTemplate{
			schema.SystemMessage(tools_control.SystemMessageNormalTemplate),
			//schema.MessagesPlaceholder("examples", true),
			schema.MessagesPlaceholder("chat_history", true),
			schema.UserMessage(tools_control.UserMessageTemplate),
		},
	}
	ctp = &ChatTemplateImpl{config: config}
	return ctp, nil
}

/*初始化*/
func (impl *ChatTemplateImpl) Format(ctx context.Context, vs map[string]any, opts ...prompt.Option) ([]*schema.Message, error) {
	/*初始化模版*/
	template := prompt.FromMessages(impl.config.FormatType, impl.config.Templates...) //修改了此处
	format, err := template.Format(ctx, vs)
	if err != nil {
		return nil, fmt.Errorf("提示工程构建失败: %w", err)
	}
	if len(format) == 0 {
		return nil, fmt.Errorf("消息格式化结果为空")
	}
	log.Println("初始化模版：", format)
	return format, nil
}
