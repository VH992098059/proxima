// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WordsDao is the data access object for table words.
type WordsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns WordsColumns // columns contains all the column names of Table for convenient usage.
}

// WordsColumns defines and stores column names for table words.
type WordsColumns struct {
	Id             string // 主键ID
	Accent         string // 单词音标
	Cloze          string // 填空格式
	ClozeData      string // 填空元数据(JSON格式)
	DeformationImg string // 变形图片路径
	ImageFile      string // 主图片路径
	MeanCn         string // 中文释义
	MeanEn         string // 英文释义
	Options        string // 选项列表(JSON数组)
	Sentence       string // 例句
	SentenceAudio  string // 例句音频路径
	SentenceTrans  string // 例句翻译
	Word           string // 单词原文
	WordAudio      string // 单词音频路径
	WordEtyma      string // 词源解析
	WordLevelId    string // 单词级别ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// wordsColumns holds the columns for table words.
var wordsColumns = WordsColumns{
	Id:             "id",
	Accent:         "accent",
	Cloze:          "cloze",
	ClozeData:      "cloze_data",
	DeformationImg: "deformation_img",
	ImageFile:      "image_file",
	MeanCn:         "mean_cn",
	MeanEn:         "mean_en",
	Options:        "options",
	Sentence:       "sentence",
	SentenceAudio:  "sentence_audio",
	SentenceTrans:  "sentence_trans",
	Word:           "word",
	WordAudio:      "word_audio",
	WordEtyma:      "word_etyma",
	WordLevelId:    "word_level_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewWordsDao creates and returns a new DAO object for table data access.
func NewWordsDao() *WordsDao {
	return &WordsDao{
		group:   "default",
		table:   "words",
		columns: wordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WordsDao) Columns() WordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
