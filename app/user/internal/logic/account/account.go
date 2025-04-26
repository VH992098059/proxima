package account

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"proxima/app/user/api/pbentity"
	"proxima/app/user/internal/dao"
	"proxima/app/user/internal/model/entity"
	"proxima/app/user/utility"
	utility3 "proxima/utility"
	"proxima/utility/consts"
	"time"
)

type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

func Register(ctx context.Context, in *pbentity.Users) (id int, err error) {
	isExistUser, err := dao.Users.Ctx(ctx).Where("username=?", in.Username).Count()
	if err != nil {
		return 0, err
	}
	if isExistUser > 0 {
		return -1, gerror.New("用户已存在")
	}
	isExistEmail, err := dao.Users.Ctx(ctx).Where("email=?", in.Email).Count()
	if err != nil {
		return 0, err
	}
	if isExistEmail > 0 {
		return -1, gerror.New("邮箱已存在")
	}
	encryptPassword, _ := utility.Encrypt(in.Password)
	getId, err := dao.Users.Ctx(ctx).Data(pbentity.Users{Username: in.Username, Email: in.Email, Password: encryptPassword}).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return int(getId), nil
}

func Login(ctx context.Context, username, password string) (token string, err error) {
	var user entity.Users
	if username != "" {
		err = dao.Users.Ctx(ctx).Fields("id", "username", "password", "email").Where("username", username).Scan(&user)
		if err != nil {
			return "", gerror.New("用户名不存在")
		}
	}
	passwordComparison := utility.Verify(password, user.Password)
	if !passwordComparison {
		return "", gerror.New("用户名或密码错误")
	}

	us := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	//JWT加密
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, us).SignedString([]byte(consts.JwtKey))
	if err != nil {
		log.Println(err)
		return "", err
	}
	//存储redis
	err = utility3.SetJWT(ctx, user.Username, signedString, 3600*24)
	if err != nil {
		log.Println("出错啦！")
		return "", err
	}
	return signedString, nil
}
