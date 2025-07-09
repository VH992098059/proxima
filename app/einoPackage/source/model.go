package source

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func newChatModel(ctx context.Context) (cm model.ChatModel, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("加载 .env 文件出错")
	}
	// 检查环境变量
	if os.Getenv("Model_Type") == "" || os.Getenv("Openai_API_Key") == "" || os.Getenv("Base_URL") == "" {
		log.Fatal(".env 未配置 Model_Type, Openai_API_Key, Base_URL")
	}
	// TODO Modify component configuration here.
	config := &openai.ChatModelConfig{
		Model:   os.Getenv("Model_Type"),
		APIKey:  os.Getenv("Openai_API_Key"),
		BaseURL: os.Getenv("Base_URL"),
	}
	cm, err = openai.NewChatModel(ctx, config)
	if err != nil {
		return nil, err
	}
	return cm, nil
}
