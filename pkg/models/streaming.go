package models

type Streaming struct {
	GormModel
	VedioSourceUrl string `gorm:"unique;type:varchar(512)"`
}
