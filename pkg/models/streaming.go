package models

type Streaming struct {
	GormModel
	VedioUrl string `gorm:"unique;type:varchar(512)"`
	//type
}
