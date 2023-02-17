package models

type UserRole int

const (
	RolePlayer    UserRole = 0
	RoleOperation UserRole = 100
	RoleAdmin     UserRole = 1000
)

//Balance  uint64 `gorm:"type:bigint;not null;default:0"` //need default whend add column and type:not null
type User struct {
	GormModel
	Name     string   `gorm:"unique;type:varchar(128)"`          //trons varchar(255)
	Account  string   `gorm:"unique;type:varchar(128);not null"` //trons varchar(255)
	Password string   `gorm:"type:char(64);not null"`            //or char(60), len of hashed password
	Role     UserRole `gorm:"not null;default:0"`
}
