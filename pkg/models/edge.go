package models

//Status int    `gorm:"type:tinyint unsigned;not null"`
type Edge struct {
	GormModel
	IP     string     `gorm:"type:char(32);not null"`
	Port   int        `gorm:"type:int unsigned;default:0"`
	Online bool       `gorm:"default:0;not null"`
	Status EdgeStatus `gorm:"type:int unsigned;default:0;not null"` //0: init
}
