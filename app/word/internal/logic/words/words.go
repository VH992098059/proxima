package words

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"proxima/app/word/api/pbentity"
	"proxima/utility"

	"go.mongodb.org/mongo-driver/bson"
)

func Create(ctx context.Context) (id uint, err error) {
	return 1, nil
}
func Get(ctx context.Context, start, limit int) (word []*pbentity.Words, err error) {
	var words []*pbentity.Words
	//使用mongoDB数据库
	mongoClient, err := utility.ConnectMongo()
	if err != nil {
		return nil, err
	}
	defer mongoClient.Disconnect(ctx)
	collection := mongoClient.Database("english").Collection("0")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 创建一个临时结构体来匹配MongoDB中的数据结构
	type MongoWord struct {
		ID            primitive.ObjectID `bson:"_id,omitempty"`
		Word          string             `bson:"word"`
		Accent        string             `bson:"accent"`
		MeanCn        string             `bson:"mean_cn"`
		SentenceTrans string             `bson:"sentence_trans"`
		Sentence      string             `bson:"sentence"`
	}
	var mongoWords []MongoWord
	if err = cursor.All(ctx, &mongoWords); err != nil {
		return nil, err
	}

	// 将MongoDB数据转换为pbentity.Words结构
	words = make([]*pbentity.Words, len(mongoWords))

	for i, w := range mongoWords {
		words[i] = &pbentity.Words{
			WordId:        w.ID.String(),
			Word:          w.Word,
			Accent:        w.Accent,
			MeanCn:        w.MeanCn,
			Sentence:      w.Sentence,
			SentenceTrans: w.SentenceTrans,
		}

	}
	// 如果获取的数量超过请求的数量
	if len(words) > limit {
		/*// 初始化随机数种子，随机选择指定数量的单词
		rand.NewSource(time.Now().UnixNano())
		// 使用Fisher-Yates洗牌算法随机打乱单词顺序
		for i := len(words) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			words[i], words[j] = words[j], words[i]
		}*/
		// 选择前number个单词
		words = words[start:limit]
	}
	//使用MySQL数据库
	/*err = dao.Words.Ctx(ctx).Fields("accent", "cloze", "mean_cn", "sentence_trans", "word").OrderRandom().Limit(number).Scan(&words)
	if err != nil {
		return nil, err
	}*/
	return words, nil
}
