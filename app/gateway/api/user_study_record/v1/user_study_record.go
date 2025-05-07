package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type InsertReq struct {
	g.Meta         `path:"words/insert_words" method:"post" sm:"添加学习过的单词" tags:"单词"`
	WordId         string      `json:"word_id" v:"required"`
	StudyTime      *gtime.Time `json:"study_time"`
	MasteryLevel   uint32      `json:"mastery_level"`
	ReviewCount    uint32      `json:"review_count"`
	LastReviewTime *gtime.Time `json:"last_review_time"`
	WordLevelId    uint32      `json:"word_level_id"`
	IsFavorite     bool        `json:"is_favorite"`
}
type InsertRes struct {
}
