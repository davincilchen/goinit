package models

//VedioUrl string `gorm:"unique;type:varchar(512)"`
type Streaming struct {
	GormModel
	VedioUrl string `gorm:"type:varchar(512)"`
	//type
}
