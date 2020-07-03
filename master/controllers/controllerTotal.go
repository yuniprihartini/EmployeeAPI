package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/master/usecase"

	"github.com/gorilla/mux"
)

type TotalHandler struct {
	TotalUseCase usecase.TotalUseCase
}

func TotalController(r *mux.Router, service usecase.TotalUseCase) {
	totalHandler := TotalHandler{service}
	r.HandleFunc("/total_gaji_perbulan", totalHandler.GetSalaryPerMonth).Methods(http.MethodGet)
	r.HandleFunc("/total_gaji_perkaryawan", totalHandler.GetSalaryPerEmployee).Methods(http.MethodGet)
	r.HandleFunc("/total_karyawan_tetap", totalHandler.GetTotalFixEmployee).Methods(http.MethodGet)
	r.HandleFunc("/total_karyawan_kontrak", totalHandler.GetTotalTemporaryEmployee).Methods(http.MethodGet)
	r.HandleFunc("/total_karyawan_aktif", totalHandler.GetAlltotalActiveEmployee).Methods(http.MethodGet)
}

func (t TotalHandler) GetTotalFixEmployee(w http.ResponseWriter, r *http.Request) {
	totals, err := t.TotalUseCase.GetTotalFixEmployee()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOftotals, _ := json.Marshal(totals)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOftotals)
}
func (t TotalHandler) GetTotalTemporaryEmployee(w http.ResponseWriter, r *http.Request) {
	totals, err := t.TotalUseCase.GetTotalTemporaryEmployee()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOftotals, _ := json.Marshal(totals)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOftotals)
}
func (t TotalHandler) GetAlltotalActiveEmployee(w http.ResponseWriter, r *http.Request) {
	totals, err := t.TotalUseCase.GetTotalActiveEmployee()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOftotals, _ := json.Marshal(totals)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOftotals)
}
func (t TotalHandler) GetSalaryPerEmployee(w http.ResponseWriter, r *http.Request) {
	totals, err := t.TotalUseCase.GetTotalSalaryPerEmployee()

	if err != nil {
		w.Write([]byte("Data not found"))
		fmt.Println(err)
	}
	byteOftotals, _ := json.Marshal(totals)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
		fmt.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOftotals)
}
func (t TotalHandler) GetSalaryPerMonth(w http.ResponseWriter, r *http.Request) {
	totals, err := t.TotalUseCase.GetTotalSalaryPerMonth()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOftotals, _ := json.Marshal(totals)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOftotals)
}
