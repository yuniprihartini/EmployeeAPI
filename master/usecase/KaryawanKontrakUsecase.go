package usecase

import "test/master/models"

type KaryawanKontrakUseCase interface {
	GetAllTemporaryEmployee() ([]*models.KaryawanKontrak, error)
	GetById(id string) (*models.KaryawanKontrak, error)
	InsertTemporaryEmployee(karyawanKontrak models.KaryawanKontrak) error
	UpdateTemporaryEmployee(id string, karyawanKontrak models.KaryawanKontrak) error
	DeleteTemporaryEmployeeById(id string) error
}
