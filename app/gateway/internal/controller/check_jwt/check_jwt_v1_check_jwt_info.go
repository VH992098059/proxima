package check_jwt

import (
	"context"
	"fmt"

	"log"
	"proxima/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "proxima/app/gateway/api/check_jwt/v1"
)

func (c *ControllerV1) CheckJwtInfo(ctx context.Context, req *v1.CheckJwtInfoReq) (res *v1.CheckJwtInfoRes, err error) {
	getJwt := utility.GetJWT(ctx)
	if getJwt == "" {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "token is empty")
		return nil, err
	}
	log.Println("用户token：", getJwt)
	// 使用自定义的JwtClaims结构体解析JWT
	claims := &utility.JwtClaims{}
	token, err := utility.Decryption(getJwt, claims)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "token无效")
	}

	userKey := fmt.Sprintf("user:%s", claims.Username)
	checkJWT, err := utility.CheckJWT(ctx, userKey, getJwt)
	if err != nil {
		return nil, err
	}
	checkJWTBlack, err := utility.CheckBlackTokens(ctx, claims.Username, getJwt)
	log.Println("验证token是否在redis黑名单：", checkJWTBlack)
	if err != nil {
		log.Fatal("check_jwt出错")
		return nil, err
	}
	if !checkJWT || checkJWTBlack {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "token已失效或不存在")
	}
	return &v1.CheckJwtInfoRes{
		Msg: "验证成功",
	}, nil
}
