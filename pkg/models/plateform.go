package models

type Plateform struct {
	GormModel
	Name  string `gorm:"type:varchar(64);not null"`
	Brief string `gorm:"type:varchar(256)"`
}
