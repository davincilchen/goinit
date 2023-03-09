package models

type EdgeApp struct {
	GormModel
	// IP     string `gorm:"type:char(32);not null"`
	// Port   int    `gorm:"type:int unsigned;default:0;not null"` //0: init
	AppID  int `gorm:"not null"`
	App    App
	EdgeID int `gorm:"not null"`
	Edge   Edge
}
