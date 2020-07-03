package repositories

import (
	"database/sql"
	"fmt"
	"test/master/models"
	"test/tools"
)

type TotalRepoImple struct {
	db *sql.DB
}

func (t *TotalRepoImple) GetTotalFixEmployee() (int, error) {
	var total int
	err := t.db.QueryRow(tools.GET_TOTAL_FIX_EMPLOYEE).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (t *TotalRepoImple) GetTotalSalaryPerEmployee() ([]models.TotalGajiPerKaryawan, error) {
	var result []models.TotalGajiPerKaryawan
	data, err := t.db.Query(tools.GET_TOTAL_SALARY_PEREMPLOYEE)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		var each models.TotalGajiPerKaryawan
		err = data.Scan(&each.Nama, &each.Gaji)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}

func (t *TotalRepoImple) GetTotalSalaryPerMonth() (int, error) {
	var total int
	err := t.db.QueryRow(tools.GET_TOTAL_SALARY_PERMONTH).Scan(&total)
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (t *TotalRepoImple) GetTotalTemporaryEmployee() (int, error) {
	var total int
	err := t.db.QueryRow(tools.GET_TOTAL_TEMPORARY_EMPLOYEE).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (t *TotalRepoImple) GetTotalActiveEmployee() (int, error) {
	var total int
	err := t.db.QueryRow(tools.GET_TOTAL_ACTIVE_EMPLOYEE).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func InitTotalRepoImpl(db *sql.DB) TotalRepository {
	return &TotalRepoImple{db}
}
