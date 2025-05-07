package user_study_record

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"proxima/utility"

	"github.com/gogf/gf/v2/os/gtime"
)

type WordProgressSimple struct {
	WordId         string      `bson:"word_id"`
	StudyTime      *gtime.Time `bson:"study_time"`
	MasteryLevel   uint32      `bson:"mastery_level"`
	ReviewCount    uint32      `bson:"review_count"`
	LastReviewTime *gtime.Time `bson:"last_review_time"`
	WordLevelId    uint32      `json:"word_level_id"`
	IsFavorite     bool        `bson:"is_favorite"`
}

// InsertRecord 添加已学习的单词
func InsertRecord(ctx context.Context, in WordProgressSimple, token string) (msg string, err error) {
	// 解析token获取用户信息
	decryption, err := utility.Decryption(token, &utility.JwtClaims{})
	if err != nil {
		return "", utility.WrapError(err, "token解析失败")
	}
	claims := decryption.Claims.(*utility.JwtClaims)
	if claims.Username == "" {
		return "", utility.WrapError(err, "用户名为空")
	}
	// 连接MongoDB数据库
	mongoClient, err := utility.ConnectMongo()
	if err != nil {
		return "", utility.WrapError(err, "MongoDB连接失败")
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			utility.LogError("MongoDB断开连接失败:", err)
		}
	}()
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"word_id", 1}},          // 以升序为 "word_id" 字段创建索引
		Options: options.Index().SetUnique(true), // 设置 Unique 选项为 true
	}
	// 获取数据库集合
	collection := mongoClient.Database("user_records").Collection(claims.Username)
	utility.LogInfo("准备插入单词记录", map[string]interface{}{
		"username":         claims.Username,
		"word_id":          in.WordId,
		"study_time":       in.StudyTime,
		"last_review_time": in.LastReviewTime,
	})
	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return "", err
	}

	// 插入数据
	_, err = collection.InsertOne(ctx, &WordProgressSimple{
		WordId:         in.WordId,
		StudyTime:      in.StudyTime,
		MasteryLevel:   in.MasteryLevel,
		ReviewCount:    in.ReviewCount,
		LastReviewTime: in.LastReviewTime,
		WordLevelId:    in.WordLevelId,
		IsFavorite:     in.IsFavorite,
	})
	if err != nil {
		return "", utility.WrapError(err, "MongoDB插入数据失败")
	}

	utility.LogInfo("单词记录插入成功", map[string]interface{}{
		"username": claims.Username,
		"word_id":  in.WordId,
	})
	return "插入成功", nil
}
