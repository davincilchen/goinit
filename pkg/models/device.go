package models

type Device struct {
	GormModel
	UUID string `gorm:"type:char(64);not null"`
	Type int    `gorm:"type:tinyint unsigned"` //0:眼鏡 1:手機
}
