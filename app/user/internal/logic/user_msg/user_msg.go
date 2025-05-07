package user_msg

import (
	"context"
	"log"
	"proxima/app/user/internal/dao"
	"proxima/app/user/internal/model/entity"
	"proxima/utility"
)

func Info(ctx context.Context, token string) (user *entity.Users, err error) {
	log.Println("Parsing JWT:", token)
	// 解析自定义 Claims
	claims := &utility.JwtClaims{}
	parsedToken, err := utility.Decryption(token, claims)
	if err != nil {
		return nil, err
	}
	log.Println(parsedToken.Claims)
	// 从 claims.Id 获取用户 ID 并查询
	userID := claims.Id
	err = dao.Users.Ctx(ctx).Fields("username", "email", "created_at", "updated_at").Where("id", userID).Scan(&user)
	return
}
