package words

import (
	"context"
	"proxima/app/word/api/pbentity"
	v1 "proxima/app/word/api/words/v1"
	"proxima/app/word/internal/logic/words"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedWordsServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterWordsServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

// Get 获取单词
func (*Controller) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	wordsList, err := words.Get(ctx, int(req.Start), int(req.Limit))
	if err != nil {
		return nil, err
	}
	pbWords := make([]*pbentity.Words, len(wordsList))
	for i, word := range wordsList {
		pbWords[i] = &pbentity.Words{
			WordId:        word.WordId,
			Word:          word.Word,
			Accent:        word.Accent,
			MeanCn:        word.MeanCn,
			Sentence:      word.Sentence,
			SentenceTrans: word.SentenceTrans,
		}
	}
	return &v1.GetRes{
		Words: pbWords,
	}, nil
}
