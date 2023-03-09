package usecase

import (
	repo "xr-central/pkg/app/app/repo/mysql"
	"xr-central/pkg/models"
)

var appGenreRepo repo.AppGenre

// ============================================= //
type AppHandle struct {
}

func (t *AppHandle) RegGenre(data *models.AppGenre) (*models.AppGenre, error) {
	return appGenreRepo.RegType(data)
}

func (t *AppHandle) GetGenre(id uint) (*models.AppGenre, error) {
	return appGenreRepo.Get(id)
}
