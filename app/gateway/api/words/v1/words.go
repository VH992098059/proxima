package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta     `path:"words/create" method:"post" sm:"创建" tags:"单词"`
	Word       string `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition string `json:"definition" v:"required|length:1,300" dc:"单词定义"`
}

type CreateRes struct {
}

type DetailReq struct {
	g.Meta `path:"words/get" method:"get" sm:"详情" tags:"单词"`
	Id     uint `json:"id" `
	Start  uint `json:"start" v:"required|min:1"`
	Limit  uint `json:"limit" v:"required|min:1"`
}
type DetailRes struct {
	WordList interface{} `json:"word_list"`
}
