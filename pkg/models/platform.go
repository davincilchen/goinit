package models

type Platform struct {
	GormModel
	Name  string `gorm:"unique;type:varchar(64);not null"`
	Brief string `gorm:"type:varchar(256)"`
}
