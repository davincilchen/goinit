package usecase

import "xr-central/pkg/models"

type Edge struct {
	models.Edge
}

func (t *Edge) Release() error {

	return nil
}

func (t *Edge) Resume() error {

	return nil
}

func (t *Edge) Status() error {

	return nil
}

func (t *Edge) StartAPP() error {

	return nil
}

func (t *Edge) StopAPP() error {

	return nil
}
