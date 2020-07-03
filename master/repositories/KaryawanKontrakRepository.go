package repositories

import "test/master/models"

type KaryawanKontrakRepository interface {
	GetAllTemporaryEmployee() ([]*models.KaryawanKontrak, error)
	GetById(id string) (*models.KaryawanKontrak, error)
	InsertTemporaryEmployee(karyawanTetap models.KaryawanKontrak) error
	UpdateTemporaryEmployee(id string, karyawanKontrak models.KaryawanKontrak) error
	DeleteTemporaryEmployeeById(id string) error
}
