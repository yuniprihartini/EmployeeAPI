package repositories

import "test/master/models"

type KaryawanTetapRepository interface {
	GetAllFixEmployee() ([]*models.KaryawanTetap, error)
	GetById(id string) (*models.KaryawanTetap, error)
	InsertFixEmployee(karyawanTetap models.KaryawanTetap) error
	UpdateFixEmployee(id string, karyawanTetap models.KaryawanTetap) error
	DeleteFixEmployeeById(id string) error
}
