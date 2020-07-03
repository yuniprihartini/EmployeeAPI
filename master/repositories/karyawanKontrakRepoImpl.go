package repositories

import (
	"database/sql"
	"fmt"
	"test/master/models"
	"test/tools"
)

// TemporaryEmployeeRepoImpl struct
type KaryawanKontrakRepoImple struct {
	db *sql.DB
}

// GetAllTemporaryEmployee method
func (kt *KaryawanKontrakRepoImple) GetAllTemporaryEmployee() ([]*models.KaryawanKontrak, error) {
	var dataKaryawanKontrak []*models.KaryawanKontrak
	data, err := kt.db.Query(tools.GET_ALL_TEMPORARY_EMPLOYEE)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		var KaryawanKontrak = models.KaryawanKontrak{}
		err := data.Scan(&KaryawanKontrak.KaryawanKontrakId, &KaryawanKontrak.NamaLengkap, &KaryawanKontrak.Alamat, &KaryawanKontrak.TempatLahir, &KaryawanKontrak.TanggalLahir, &KaryawanKontrak.GajiPokok, &KaryawanKontrak.StatusKaryawan)
		if err != nil {
			return nil, err
		}
		dataKaryawanKontrak = append(dataKaryawanKontrak, &KaryawanKontrak)
	}
	return dataKaryawanKontrak, nil
}

// GetTemporaryEmployeeById method
func (kt *KaryawanKontrakRepoImple) GetById(id string) (*models.KaryawanKontrak, error) {
	var dataKaryawanKontrak = new(models.KaryawanKontrak)
	err := kt.db.QueryRow(tools.GET_TEMPORARY_EMPLOYEE_BY_ID, id).Scan(&dataKaryawanKontrak.KaryawanKontrakId,
		&dataKaryawanKontrak.NamaLengkap, &dataKaryawanKontrak.Alamat, &dataKaryawanKontrak.TempatLahir,
		&dataKaryawanKontrak.TanggalLahir, &dataKaryawanKontrak.GajiPokok, &dataKaryawanKontrak.StatusKaryawan)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return dataKaryawanKontrak, nil
}

// UpdateTemporaryEmployee method
func (kt *KaryawanKontrakRepoImple) UpdateTemporaryEmployee(id string, KaryawanKontrak models.KaryawanKontrak) error {
	tx, err := kt.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.UPDATE_TEMPORARY_EMPLOYEE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&KaryawanKontrak.NamaLengkap, &KaryawanKontrak.Alamat, &KaryawanKontrak.TempatLahir, &KaryawanKontrak.TanggalLahir, &KaryawanKontrak.GajiPokok, &KaryawanKontrak.StatusKaryawan, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// InsertTemporaryEmployee method
func (kt *KaryawanKontrakRepoImple) InsertTemporaryEmployee(KaryawanKontrak models.KaryawanKontrak) error {
	tx, err := kt.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.INSERT_TEMPORARY_EMPLOYEE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&KaryawanKontrak.KaryawanKontrakId, &KaryawanKontrak.NamaLengkap, &KaryawanKontrak.Alamat, &KaryawanKontrak.TempatLahir, &KaryawanKontrak.TanggalLahir, &KaryawanKontrak.GajiPokok, &KaryawanKontrak.StatusKaryawan)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// DeleteTemporaryEmployee method
func (kt *KaryawanKontrakRepoImple) DeleteTemporaryEmployeeById(id string) error {
	tx, err := kt.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.DELETE_TEMPORARY_EMPLOYEE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// TemporaryEmployeeRepoImpl init
func InitTemporaryEmployeeRepoImpl(db *sql.DB) KaryawanKontrakRepository {
	return &KaryawanKontrakRepoImple{db}
}
