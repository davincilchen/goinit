package models

type EdgeApp struct {
	GormModel
	// IP     string `gorm:"type:char(32);not null"`
	// Port   int    `gorm:"type:int unsigned;default:0;not null"` //0: init
	EdgeID uint `gorm:"not null;index:edge_app,unique"`
	Edge   Edge
	AppID  uint `gorm:"not null;index:edge_app,unique"`
	App    App
}
