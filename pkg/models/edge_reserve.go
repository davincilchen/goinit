package models

//Status int    `gorm:"type:tinyint unsigned;not null"`
type EdgeReserve struct {
	GormModel
	IP       string `gorm:"type:char(32);not null"`
	Status   int    `gorm:"type:int unsigned;default:0;not null"` //0: init
	EdgeID   int    `gorm:"not null"`
	Edge     Edge
	DeviceID int `gorm:"not null"`
	Device   Device
	AppID    int `gorm:"not null"`
	App      App
}
