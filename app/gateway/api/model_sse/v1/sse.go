package v1

import "github.com/gogf/gf/v2/frame/g"

type ModelReq struct {
	g.Meta   `path:"/ai/chat" method:"get"`
	ClientId string
	Message  string
}
type ModelRes struct {
	Message string
}
