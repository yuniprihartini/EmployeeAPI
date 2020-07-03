package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test/master/models"
	"test/master/usecase"
	"test/master/utils"
	"test/tools"

	"github.com/gorilla/mux"
)

type KaryawanKontrakHandler struct {
	KaryawanKontrakUseCase usecase.KaryawanKontrakUseCase
}

func KaryawanKontrakControllers(r *mux.Router, service usecase.KaryawanKontrakUseCase) {
	KaryawanKontrakHandler := KaryawanKontrakHandler{service}
	r.HandleFunc("/KaryawanKontraks", KaryawanKontrakHandler.GetAllKaryawanKontraks).Methods(http.MethodGet)
	r.HandleFunc("/KaryawanKontrak/{id}", KaryawanKontrakHandler.KaryawanKontrakById).Methods(http.MethodGet)
	r.HandleFunc("/KaryawanKontrak", KaryawanKontrakHandler.CreateKaryawanKontrak).Methods(http.MethodPost)
	r.HandleFunc("/KaryawanKontrak/{id}", KaryawanKontrakHandler.DeleteKaryawanKontrak).Methods(http.MethodDelete)
	r.HandleFunc("/KaryawanKontrak/{id}", KaryawanKontrakHandler.UpdateKaryawanKontrak).Methods(http.MethodPut)
}

func (kt KaryawanKontrakHandler) GetAllKaryawanKontraks(w http.ResponseWriter, r *http.Request) {
	KaryawanKontraks, err := kt.KaryawanKontrakUseCase.GetAllTemporaryEmployee()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfKaryawanKontraks, _ := json.Marshal(KaryawanKontraks)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfKaryawanKontraks)
}

func (kt KaryawanKontrakHandler) KaryawanKontrakById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	fmt.Println(param["id"])
	KaryawanKontrak, err := kt.KaryawanKontrakUseCase.GetById(param["id"])
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfKaryawanKontraks, _ := json.Marshal(KaryawanKontrak)
	if err != nil {
		w.Write([]byte("Oops, somethin when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfKaryawanKontraks)
}

func (p KaryawanKontrakHandler) CreateKaryawanKontrak(w http.ResponseWriter, r *http.Request) {
	var KaryawanKontrak models.KaryawanKontrak
	err := json.NewDecoder(r.Body).Decode(&KaryawanKontrak)

	var isTrue = utils.ValidationKaryawanKontrak(&KaryawanKontrak)
	tools.PrintlnErr(err)
	if isTrue == false {
		w.Write([]byte("Name or address must be fill"))
	} else {
		err = p.KaryawanKontrakUseCase.InsertTemporaryEmployee(KaryawanKontrak)
		tools.PrintlnErr(err)
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}

}

func (p KaryawanKontrakHandler) DeleteKaryawanKontrak(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	err := p.KaryawanKontrakUseCase.DeleteTemporaryEmployeeById(param["id"])
	tools.PrintlnErr(err)
	log.Println("Delete successful")
	w.Write([]byte("Delete successful"))
}

func (p KaryawanKontrakHandler) UpdateKaryawanKontrak(w http.ResponseWriter, r *http.Request) {
	var KaryawanKontrak models.KaryawanKontrak
	param := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&KaryawanKontrak)
	var isTrue = utils.ValidationKaryawanKontrak(&KaryawanKontrak)
	if isTrue == false {
		w.Write([]byte("Name or address must be fill"))
	} else {
		err := p.KaryawanKontrakUseCase.UpdateTemporaryEmployee(param["id"], KaryawanKontrak)
		tools.PrintlnErr(err)
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
