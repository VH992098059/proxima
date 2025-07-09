package main

import (
	"context"
	"fmt"
	"proxima/app/einoPackage/api"
)

func main() {
	ctx := context.Background()
	var input string
	for {
		fmt.Print("请输入内容：")
		fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("再见")
			return
		}
		api.ChatAiModel(ctx, true, input)
		fmt.Println()
	}

}
