package utils

import "test/master/models"

func ValidationKaryawanKontrak(input *models.KaryawanKontrak) bool {
	var result bool
	result = true
	switch {
	case input.NamaLengkap == "":
		result = false
	case input.Alamat == "":
		result = false
	}
	return result
}

func ValidationKaryawanTetap(input *models.KaryawanTetap) bool {
	var result bool
	result = true
	switch {
	case input.NamaLengkap == "":
		result = false
	case input.Alamat == "":
		result = false
	}
	return result
}
