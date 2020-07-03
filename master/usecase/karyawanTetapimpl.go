package usecase

import (
	"test/master/models"
	"test/master/repositories"
	"test/tools"
)

type KaryawanTetapUsecaseImpl struct {
	karyawanTetapRepo repositories.KaryawanTetapRepository
}

func (kt KaryawanTetapUsecaseImpl) GetAllFixEmployee() ([]*models.KaryawanTetap, error) {
	karyawanTetap, err := kt.karyawanTetapRepo.GetAllFixEmployee()
	if err != nil {
		return nil, err
	}

	return karyawanTetap, nil
}

func (kt KaryawanTetapUsecaseImpl) GetById(id string) (*models.KaryawanTetap, error) {
	karyawanTetap, err := kt.karyawanTetapRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return karyawanTetap, nil
}
func (kt KaryawanTetapUsecaseImpl) InsertFixEmployee(karyawanTetap models.KaryawanTetap) error {
	err := kt.karyawanTetapRepo.InsertFixEmployee(karyawanTetap)
	tools.FatalErr(err)

	return nil
}
func (kt KaryawanTetapUsecaseImpl) UpdateFixEmployee(id string, karyawanTetap models.KaryawanTetap) error {
	err := kt.karyawanTetapRepo.UpdateFixEmployee(id, karyawanTetap)
	tools.FatalErr(err)

	return nil
}
func (kt KaryawanTetapUsecaseImpl) DeleteFixEmployeeById(id string) error {
	err := kt.karyawanTetapRepo.DeleteFixEmployeeById(id)
	tools.FatalErr(err)

	return nil
}
func InitFixEmployeeUseCaseImpl(karyawanTetapRepo repositories.KaryawanTetapRepository) KaryawanTetapUseCase {
	return &KaryawanTetapUsecaseImpl{karyawanTetapRepo}
}
