//package repo
package mysql

import (
	"xr-central/pkg/db"
	"xr-central/pkg/models"

	"gorm.io/gorm"
)

type Edge struct {
}

func GetDB() *gorm.DB {
	return db.MainDB
}

func (t *Edge) LoadEdges() ([]models.Edge, error) {
	ddb := GetDB()
	out := []models.Edge{}

	dbc := ddb.Find(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}

func (t *Edge) FindEdgesWithAppID(appID int) ([]models.EdgeApp, error) {
	ddb := GetDB()
	out := []models.EdgeApp{}

	dbc := ddb.Where("app_id = ?", appID).Find(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil

}
