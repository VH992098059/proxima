package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeSuccess      = gcode.New(0, "成功", nil)
	CodeUserNotLogin = gcode.New(10001, "用户未登录", nil)
	CodeUserNotFound = gcode.New(10002, "用户未找到", nil)
)
