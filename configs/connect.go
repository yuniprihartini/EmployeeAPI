package configs

import (
	"database/sql"
	"test/tools"
	_"github.com/go-sql-driver/mysql"
)

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Yunip12345#@tcp(localhost:3306)/karyawan")
	err = db.Ping()
	tools.FatalErr(err)
	return db, nil
}
