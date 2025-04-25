package middleware

import (
	"fmt"
	"net/http"
	"proxima/utility/consts"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

// Auth jwt中间件
func Auth(r *ghttp.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
		return
	}
	//分割token
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
		return
	}
	//分割后的token
	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(consts.JwtKey), nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
		return
	}
	r.Middleware.Next()
}
