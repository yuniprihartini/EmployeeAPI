package models

type KaryawanTetap struct {
	KartapId       string `json:"kartapid"`
	NamaLengkap    string `json:"nama"`
	Alamat         string `json:"alamat"`
	TempatLahir    string `json:"tempatlahir"`
	TanggalLahir   string `json:"tanggallahir"`
	GajiPokok      int    `json:"gaji"`
	Tunjangan      int    `json:"tunjangan"`
	StatusKaryawan string `json:"status"`
}
