package usecase

import (
	"test/master/models"
	"test/master/repositories"
)

type TotalUseCaseImpl struct {
	TotalRepo repositories.TotalRepository
}

func (t TotalUseCaseImpl) GetTotalFixEmployee() (int, error) {
	total, err := t.TotalRepo.GetTotalFixEmployee()
	if err != nil {
		return 0, err
	}

	return total, nil
}
func (t TotalUseCaseImpl) GetTotalTemporaryEmployee() (int, error) {
	total, err := t.TotalRepo.GetTotalTemporaryEmployee()
	if err != nil {
		return 0, err
	}

	return total, nil
}
func (t TotalUseCaseImpl) GetTotalActiveEmployee() (int, error) {
	total, err := t.TotalRepo.GetTotalActiveEmployee()
	if err != nil {
		return 0, err
	}

	return total, nil
}
func (t TotalUseCaseImpl) GetTotalSalaryPerEmployee() ([]models.TotalGajiPerKaryawan, error) {
	total, err := t.TotalRepo.GetTotalSalaryPerEmployee()
	if err != nil {
		return nil, err
	}

	return total, nil
}
func (t TotalUseCaseImpl) GetTotalSalaryPerMonth() (int, error) {
	total, err := t.TotalRepo.GetTotalSalaryPerMonth()
	if err != nil {
		return 0, err
	}

	return total, nil
}
func InitTotalUseCaseImpl(totalRepo repositories.TotalRepository) TotalUseCase {
	return &TotalUseCaseImpl{totalRepo}
}
