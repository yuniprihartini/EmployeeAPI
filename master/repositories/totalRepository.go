package repositories

import "test/master/models"

type TotalRepository interface {
	GetTotalFixEmployee() (int, error)
	GetTotalTemporaryEmployee() (int, error)
	GetTotalActiveEmployee() (int, error)
	GetTotalSalaryPerEmployee() ([]models.TotalGajiPerKaryawan, error)
	GetTotalSalaryPerMonth() (int, error)
}
