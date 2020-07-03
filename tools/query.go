package tools

const (

	// fix employee
	GET_ALL_FIX_EMPLOYEE   = "SELECT * FROM kartap"
	GET_FIX_EMPLOYEE_BY_ID = "SELECT * FROM kartap WHERE kartap_id = ?"
	INSERT_FIX_EMPLOYEE    = "INSERT INTO kartap VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	UPDATE_FIX_EMPLOYEE    = `UPDATE kartap SET nama_lengkap = ?, alamat = ?, tempat_lahir = ?, 
	tanggal_lahir = ?, gaji_pokok = ?,  tunjangan = ?, status_karyawan = ? WHERE kartap_id = ?`
	DELETE_FIX_EMPLOYEE = "DELETE FROM kartap WHERE kartap_id = ?"

	// temporary employee
	GET_ALL_TEMPORARY_EMPLOYEE   = "SELECT * FROM karkontrak"
	GET_TEMPORARY_EMPLOYEE_BY_ID = "SELECT * FROM karkontrak WHERE karkontrak_id = ?"
	INSERT_TEMPORARY_EMPLOYEE    = "INSERT INTO karkontrak VALUES (?, ?, ?, ?, ?, ?, ?)"
	UPDATE_TEMPORARY_EMPLOYEE    = `UPDATE karkontrak SET nama_lengkap = ?, alamat = ?, tempat_lahir = ?, 
	tanggal_lahir = ?, gaji_pokok = ?, status_karyawan = ? WHERE karkontrak_id = ?`
	DELETE_TEMPORARY_EMPLOYEE = "DELETE FROM karkontrak WHERE karkontrak_id = ?"

	GET_TOTAL_SALARY_PERMONTH = `select(select sum(kartap.gaji_pokok + kartap.tunjangan) from kartap)
	+ (select sum(karkontrak.gaji_pokok) from karkontrak) as "Total Gaji PerBulan"`
	GET_TOTAL_SALARY_PEREMPLOYEE = ` select kartap.nama_lengkap, (kartap.gaji_pokok + kartap.tunjangan) as "total gaji" from kartap
	union select karkontrak.nama_lengkap, (karkontrak.gaji_pokok) "total gaji" from karkontrak`
	GET_TOTAL_FIX_EMPLOYEE       = `select count(*)as "total karyawan tetap"from kartap`
	GET_TOTAL_TEMPORARY_EMPLOYEE = `select count(*)as "total karyawan kontrak"from karkontrak`
	GET_TOTAL_ACTIVE_EMPLOYEE    = `select (select count(*)from kartap where status_karyawan="aktif") + (select count(*)from karkontrak where status_karyawan = "aktif") as "total karyawan aktif"`
)
