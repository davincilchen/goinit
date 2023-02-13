package models

type Edge struct {
	GormModel
	IP     string `gorm:"type:char(32);not null"`
	Online bool   `gorm:"default:0;not null"`
	Status int    `gorm:"default:0;not null"` //0: init
	//Status int    `gorm:"type:tinyint unsigned;not null"`
}
