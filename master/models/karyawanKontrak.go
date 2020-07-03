package models

type KaryawanKontrak struct {
	KaryawanKontrakId string `json:"kontrakid"`
	NamaLengkap       string `json:"nama"`
	Alamat            string `json:"alamat"`
	TempatLahir       string `json:"tempatlahir"`
	TanggalLahir      string `json:"tanggallahir"`
	GajiPokok         int    `json:"gaji"`
	StatusKaryawan    string `json:"status"`
}
