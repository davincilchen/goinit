package models

type APP struct {
	GormModel
	PlatformAppID     string `gorm:"type:varchar(32);not null"`
	PlatformID        string `gorm:"not null"`
	AppTitle          string `gorm:"type:varchar(32);not null"`
	AppGenre          string `gorm:"type:varchar(32)"` //AppGenreID
	AppBrief          string `gorm:"type:varchar(256)"`
	ImageURL          string `gorm:"type:varchar(256)"`
	Developler        string `gorm:"type:varchar(128)"`
	PublicationStatus int    `gorm:"tinyint unsigned;default:0"` //0:public 1:Private
	PkgURL            string `gorm:"type:varchar(256)"`
	ExeName           string `gorm:"type:varchar(256);not null"`
}
