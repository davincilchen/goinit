package models

//PkgURL            string `gorm:"type:varchar(256)"`
type AppGenere struct {
	GormModel
	Type  string `gorm:"type:varchar(32)"`
	Brief string `gorm:"type:varchar(256)"`
}
