package utility

import (
	"context"
	"github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"log"
	"time"
)

var (
	config = gredis.Config{
		Address: "127.0.0.1:16379",
		Db:      1,
		Pass:    "password",
	}
	ctxRedis = gctx.New()
	group    = "cache"
)

func init() {
	adapter := redis.New(&config)         // 从配置加载客户端
	operation := adapter.AdapterOperation // 全局注册适配器
	_, err := operation.Conn(ctxRedis)
	if err != nil {
		return
	}
}

// SetJWT 将 token 存入 Redis 白名单，expiration 与 JWT 自身过期时间保持一致
func SetJWT(ctx context.Context, key, token string, expiration time.Duration) error {
	err := g.Redis().SetEX(ctx, key, token, int64(expiration))
	if err != nil {
		log.Println("报错")
	}
	return err
}

// CheckJWT 检查 token 是否在白名单中
/*func CheckJWT(ctx context.Context, token string) (bool, error) {
	value, err := g.Redis().Get(ctx, key)
	if err != nil {
		// key 不存在或其他错误
		return false, err
	}
	return value.String() == token, nil
}
*/
// DeleteJWT 从白名单移除 token
func DeleteJWT(ctx context.Context, key, token string) error {
	del, err := g.Redis().Del(ctx, key)
	if err != nil {
		return err
	}
	log.Println(del)
	return nil
}
