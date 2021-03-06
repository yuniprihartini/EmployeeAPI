package usecase

import "test/master/models"

type TotalUseCase interface {
	GetTotalFixEmployee() (int, error)
	GetTotalTemporaryEmployee() (int, error)
	GetTotalActiveEmployee() (int, error)
	GetTotalSalaryPerEmployee() ([]models.TotalGajiPerKaryawan, error)
	GetTotalSalaryPerMonth() (int, error)
}
