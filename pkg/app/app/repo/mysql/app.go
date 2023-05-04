//package repo
package mysql

import (
	"goinit/pkg/db"
	"goinit/pkg/models"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return db.MainDB
}

type AppGenre struct {
}

func (t *AppGenre) RegType(dev *models.AppGenre) (*models.AppGenre, error) {
	ddb := GetDB()
	out := &models.AppGenre{}
	*out = *dev

	dbc := ddb.FirstOrCreate(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return out, nil

}

func (t *AppGenre) Get(id uint) (*models.AppGenre, error) {

	ddb := GetDB()
	out := &models.AppGenre{}
	out.ID = id
	dbc := ddb.
		First(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}

type App struct {
}

func (t *App) GetApps(valid bool) ([]*models.App, error) {
	ddb := GetDB()
	out := []*models.App{}

	dbc := ddb.Preload("AppGenre").Where("valid = ?", valid).Find(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}
