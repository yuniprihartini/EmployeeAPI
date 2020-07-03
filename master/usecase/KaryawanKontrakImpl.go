package usecase

import (
	"test/master/models"
	"test/master/repositories"

	"test/tools"
)

type KaryawanKontrakUsecaseImpl struct {
	karyawanKontrakRepo repositories.KaryawanKontrakRepository
}

func (kt KaryawanKontrakUsecaseImpl) GetAllTemporaryEmployee() ([]*models.KaryawanKontrak, error) {
	karyawanKontrak, err := kt.karyawanKontrakRepo.GetAllTemporaryEmployee()
	if err != nil {
		return nil, err
	}

	return karyawanKontrak, nil
}

func (kt KaryawanKontrakUsecaseImpl) GetById(id string) (*models.KaryawanKontrak, error) {
	karyawanKontrak, err := kt.karyawanKontrakRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return karyawanKontrak, nil
}
func (kt KaryawanKontrakUsecaseImpl) InsertTemporaryEmployee(karyawanKontrak models.KaryawanKontrak) error {
	err := kt.karyawanKontrakRepo.InsertTemporaryEmployee(karyawanKontrak)
	tools.FatalErr(err)

	return nil
}
func (kt KaryawanKontrakUsecaseImpl) UpdateTemporaryEmployee(id string, karyawanKontrak models.KaryawanKontrak) error {
	err := kt.karyawanKontrakRepo.UpdateTemporaryEmployee(id, karyawanKontrak)
	tools.FatalErr(err)

	return nil
}
func (kt KaryawanKontrakUsecaseImpl) DeleteTemporaryEmployeeById(id string) error {
	err := kt.karyawanKontrakRepo.DeleteTemporaryEmployeeById(id)
	tools.FatalErr(err)

	return nil
}
func InitTemporaryEmployeeUseCaseImpl(karyawanKontrakRepo repositories.KaryawanKontrakRepository) KaryawanKontrakUseCase {
	return &KaryawanKontrakUsecaseImpl{karyawanKontrakRepo}
}
