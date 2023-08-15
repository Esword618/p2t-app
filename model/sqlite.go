package model

import "gorm.io/gorm"

type ApiType string

const (
	Simpletex ApiType = "simpletex"
	Pix2text  ApiType = "pix2text"
)

// LatexModel 代表数据库中的一个LaTeX模型。
type LatexModel struct {
	gorm.Model         // 提供了ID, CreatedAt, UpdatedAt, DeletedAt这些字段。
	Uuid       string  `gorm:"type:varchar(255);unique"` // Uuid被用作唯一识别符。
	ApiType    ApiType `gorm:"type:varchar(512)"`        // 识别使用api siplextex or pix2text
	Latex      string  `gorm:"type:text"`                // Latex表示了一个可变长度的数学公式。
	Image      string  `gorm:"type:varchar(512)"`        // Image存储了相关图片的本地路径。
}
