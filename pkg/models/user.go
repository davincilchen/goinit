package models

type User struct {
	GormModel
	Name     string `gorm:"unique;type:varchar(128)"` //trons varchar(255)
	Password string `gorm:"type:char(64)"`            //or char(60), len of hashed password
	//Balance  uint64 `gorm:"type:bigint;not null;default:0"` //need default whend add column and type:not null
}
