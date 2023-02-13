package models

type Device struct {
	GormModel
	UUID string `gorm:"type:char(64)"`
	Type int    `gorm:"type:tinyint unsigned;not null"` //0:眼鏡 1:手機
	// IP     string `gorm:"type:char(32);not null"`
	// Online bool   `gorm:"default:0;not null"`
	// Status int    `gorm:"default:0;not null"` //0: init
	//Status int    `gorm:"type:tinyint unsigned;not null"`
}
