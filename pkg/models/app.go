package models

//PkgURL            string `gorm:"type:varchar(256)"`
type App struct {
	//gorm.Model
	GormModel
	PlatformAppID     string `gorm:"type:varchar(32)"` //not null
	PlatformID        int    `gorm:"not null"`
	Platform          Platform
	AppTitle          string `gorm:"type:varchar(32);not null"`
	AppBrief          string `gorm:"type:varchar(256)"`
	ImageURL          string `gorm:"type:varchar(256)"`
	Developler        string `gorm:"type:varchar(128)"`
	PublicationStatus int    `gorm:"tinyint unsigned;default:0"` //0:public 1:Private
	APPGenreID        int    `gorm:"default:1;not null"`
	AppGenre          AppGenre
	SouceType         int    `gorm:"int unsigned;default:0"`      //0: none //EX:google drive
	Type              int    `gorm:"type:int unsigned;default:0"` //0: none //TODO: remove
	ExeName           string `gorm:"type:varchar(64);not null"`
}
