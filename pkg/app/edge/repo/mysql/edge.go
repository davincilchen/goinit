//package repo
package mysql

import (
	"fmt"
	"xr-central/pkg/db"
	"xr-central/pkg/models"

	"gorm.io/gorm"
)

type Edge struct {
}

func GetDB() *gorm.DB {
	return db.MainDB
}

func (t *Edge) GetEdges() ([]models.Edge, error) {
	ddb := GetDB()
	out := []models.Edge{}

	dbc := ddb.Find(&out)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return out, nil
}

func (t *Edge) RegEdge(ip string, port int) (*models.Edge, error) {
	ddb := GetDB()
	out := models.Edge{
		IP:   ip,
		Port: port,
	}

	dbc := ddb.FirstOrCreate(&out,
		"ip = ? AND port = ?", ip, port)

	if dbc.Error != nil {
		return nil, dbc.Error
	}

	out.Online = true
	ddb.Model(&out).Updates(models.Edge{Online: out.Online})
	return &out, nil
}

func (t *Edge) RegApps(edgeID uint, appsID []uint) ([]models.EdgeApp, error) {
	ddb := GetDB()

	var ret []models.EdgeApp
	for _, v := range appsID {
		eApp := models.EdgeApp{
			EdgeID: edgeID,
			AppID:  v,
		}
		dbc := ddb.Where(models.EdgeApp{EdgeID: edgeID,
			AppID: v}).
			FirstOrCreate(&eApp)

		if dbc.Error != nil {
			continue
		}

		fmt.Printf("[REG APP] ID = %d eApp.EdgeID = %d  eApp.AppID =%d \n", eApp.ID, eApp.EdgeID, eApp.AppID)
		ret = append(ret, eApp)
	}

	return ret, nil
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
