package models

//Status int    `gorm:"type:tinyint unsigned;not null"`
type EdgeOrder struct {
	GormModel
	IP       string `gorm:"type:char(32);not null"`
	Status   int    `gorm:"default:0;not null"` //0: init
	EdgeID   int
	Edge     Edge
	DeviceID int
	Device   Device
	AppID    int
}
