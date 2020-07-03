package master

import (
	"database/sql"
	"test/master/controllers"
	"test/master/repositories"
	"test/master/usecase"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	// Karyawantetap
	karyawanTetapRepo := repositories.InitFixEmployeeRepoImpl(db)
	karyawanTetapUseCase := usecase.InitFixEmployeeUseCaseImpl(karyawanTetapRepo)
	controllers.KaryawanTetapControllers(r, karyawanTetapUseCase)

	// KaryawanKontrak
	karyawanKontrakRepo := repositories.InitTemporaryEmployeeRepoImpl(db)
	karyawanKontrakUseCase := usecase.InitTemporaryEmployeeUseCaseImpl(karyawanKontrakRepo)
	controllers.KaryawanKontrakControllers(r, karyawanKontrakUseCase)

	// KaryawanKontrak
	totalRepo := repositories.InitTotalRepoImpl(db)
	totalUseCase := usecase.InitTotalUseCaseImpl(totalRepo)
	controllers.TotalController(r, totalUseCase)
}
