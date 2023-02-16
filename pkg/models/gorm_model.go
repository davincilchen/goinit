package models

import (
	"time"

	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

// type GormModel struct {
// 	gorm.Model
// }

type GormModel struct {
	Model
}

//CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
type Model struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t *GormModel) GetID() uint {
	return t.ID
}

func (t *GormModel) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *GormModel) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

//func (t *GormModel) GetDeletedAt() *time.Time {
func (t *GormModel) GetDeletedAt() gorm.DeletedAt {
	return t.DeletedAt
}
