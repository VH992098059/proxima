// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure of table words for DAO operations like Where/Data.
type Words struct {
	g.Meta         `orm:"table:words, do:true"`
	Id             interface{} // 主键ID
	Accent         interface{} // 单词音标
	Cloze          interface{} // 填空格式
	ClozeData      interface{} // 填空元数据(JSON格式)
	DeformationImg interface{} // 变形图片路径
	ImageFile      interface{} // 主图片路径
	MeanCn         interface{} // 中文释义
	MeanEn         interface{} // 英文释义
	Options        interface{} // 选项列表(JSON数组)
	Sentence       interface{} // 例句
	SentenceAudio  interface{} // 例句音频路径
	SentenceTrans  interface{} // 例句翻译
	Word           interface{} // 单词原文
	WordAudio      interface{} // 单词音频路径
	WordEtyma      interface{} // 词源解析
	WordLevelId    interface{} // 单词级别ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
