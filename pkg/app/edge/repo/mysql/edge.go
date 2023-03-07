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
