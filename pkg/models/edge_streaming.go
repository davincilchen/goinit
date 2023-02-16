package models

//Status   int    `gorm:"type:tinyint unsigned;not null"`
type EdgeStreaming struct {
	GormModel
	EdgeID   uint `gorm:"not null"`
	Edge     Edge
	VedioUrl string `gorm:"type:varchar(512)"`
	Status   int    `gorm:"type:tinyint unsigned;default:0;not null"`
}
