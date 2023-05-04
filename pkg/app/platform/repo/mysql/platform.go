package mysql

import (
	"initpkg/pkg/db"
	"initpkg/pkg/models"

	"gorm.io/gorm"
)

type Platform struct {
}

func GetDB() *gorm.DB {
	return db.MainDB
}
func (t *Platform) CreatePlatform(in *models.Platform) (*models.Platform, error) {
	ddb := GetDB()
	out := &models.Platform{}
	*out = *in

	dbc := ddb.Where("name = ?",
		out.Name).
		FirstOrCreate(&out) //TODO:  return error if user is exist

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}

func (t *Platform) GetPlatform(id int) (*models.Platform, error) {
	ddb := GetDB()
	out := &models.Platform{}

	dbc := ddb.Where("id = ?",
		id).
		First(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}
