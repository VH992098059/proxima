// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure for table words.
type Words struct {
	Id             uint        `json:"id"             orm:"id"              description:"主键ID"`          // 主键ID
	Accent         string      `json:"accent"         orm:"accent"          description:"单词音标"`          // 单词音标
	Cloze          string      `json:"cloze"          orm:"cloze"           description:"填空格式"`          // 填空格式
	ClozeData      string      `json:"clozeData"      orm:"cloze_data"      description:"填空元数据(JSON格式)"` // 填空元数据(JSON格式)
	DeformationImg string      `json:"deformationImg" orm:"deformation_img" description:"变形图片路径"`        // 变形图片路径
	ImageFile      string      `json:"imageFile"      orm:"image_file"      description:"主图片路径"`         // 主图片路径
	MeanCn         string      `json:"meanCn"         orm:"mean_cn"         description:"中文释义"`          // 中文释义
	MeanEn         string      `json:"meanEn"         orm:"mean_en"         description:"英文释义"`          // 英文释义
	Options        string      `json:"options"        orm:"options"         description:"选项列表(JSON数组)"`  // 选项列表(JSON数组)
	Sentence       string      `json:"sentence"       orm:"sentence"        description:"例句"`            // 例句
	SentenceAudio  string      `json:"sentenceAudio"  orm:"sentence_audio"  description:"例句音频路径"`        // 例句音频路径
	SentenceTrans  string      `json:"sentenceTrans"  orm:"sentence_trans"  description:"例句翻译"`          // 例句翻译
	Word           string      `json:"word"           orm:"word"            description:"单词原文"`          // 单词原文
	WordAudio      string      `json:"wordAudio"      orm:"word_audio"      description:"单词音频路径"`        // 单词音频路径
	WordEtyma      string      `json:"wordEtyma"      orm:"word_etyma"      description:"词源解析"`          // 词源解析
	WordLevelId    string      `json:"wordLevelId"    orm:"word_level_id"   description:"单词级别ID"`        // 单词级别ID
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`          // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`          // 更新时间
}
