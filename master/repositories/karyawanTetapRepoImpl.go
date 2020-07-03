package repositories

import (
	"database/sql"
	"fmt"
	"test/master/models"
	"test/tools"
)

// FixEmployeeRepoImpl struct
type KaryawanTetapRepoImple struct {
	db *sql.DB
}

// GetAllFixEmployee method
func (kt *KaryawanTetapRepoImple) GetAllFixEmployee() ([]*models.KaryawanTetap, error) {
	var dataKaryawanTetap []*models.KaryawanTetap
	data, err := kt.db.Query(tools.GET_ALL_FIX_EMPLOYEE)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		var karyawanTetap = models.KaryawanTetap{}
		err := data.Scan(&karyawanTetap.KartapId, &karyawanTetap.NamaLengkap, &karyawanTetap.Alamat, &karyawanTetap.TempatLahir, &karyawanTetap.TanggalLahir, &karyawanTetap.GajiPokok, &karyawanTetap.Tunjangan, &karyawanTetap.StatusKaryawan)
		if err != nil {
			return nil, err
		}
		dataKaryawanTetap = append(dataKaryawanTetap, &karyawanTetap)
	}
	return dataKaryawanTetap, nil
}

// GetFixEmployeeById method
func (kt *KaryawanTetapRepoImple) GetById(id string) (*models.KaryawanTetap, error) {
	var dataKaryawanTetap = new(models.KaryawanTetap)
	err := kt.db.QueryRow(tools.GET_FIX_EMPLOYEE_BY_ID, id).Scan(&dataKaryawanTetap.KartapId,
		&dataKaryawanTetap.NamaLengkap, &dataKaryawanTetap.Alamat, &dataKaryawanTetap.TempatLahir,
		&dataKaryawanTetap.TanggalLahir, &dataKaryawanTetap.GajiPokok, &dataKaryawanTetap.Tunjangan,
		&dataKaryawanTetap.StatusKaryawan)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return dataKaryawanTetap, nil
}

// UpdateFixEmployee method
func (kt *KaryawanTetapRepoImple) UpdateFixEmployee(id string, karyawanTetap models.KaryawanTetap) error {
	tx, err := kt.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.UPDATE_FIX_EMPLOYEE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&karyawanTetap.NamaLengkap, &karyawanTetap.Alamat, &karyawanTetap.TempatLahir, &karyawanTetap.TanggalLahir, &karyawanTetap.GajiPokok, &karyawanTetap.Tunjangan, &karyawanTetap.StatusKaryawan, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// InsertFixEmployee method
func (kt *KaryawanTetapRepoImple) InsertFixEmployee(karyawanTetap models.KaryawanTetap) error {
	tx, err := kt.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.INSERT_FIX_EMPLOYEE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&karyawanTetap.KartapId, &karyawanTetap.NamaLengkap, &karyawanTetap.Alamat, &karyawanTetap.TempatLahir, &karyawanTetap.TanggalLahir, &karyawanTetap.GajiPokok, &karyawanTetap.Tunjangan, &karyawanTetap.StatusKaryawan)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// DeleteFixEmployee method
func (kt *KaryawanTetapRepoImple) DeleteFixEmployeeById(id string) error {
	tx, err := kt.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.DELETE_FIX_EMPLOYEE)
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

// FixEmployeeRepoImpl init
func InitFixEmployeeRepoImpl(db *sql.DB) KaryawanTetapRepository {
	return &KaryawanTetapRepoImple{db}
}
