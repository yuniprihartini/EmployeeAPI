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

type KaryawanTetapHandler struct {
	KaryawanTetapUseCase usecase.KaryawanTetapUseCase
}

func KaryawanTetapControllers(r *mux.Router, service usecase.KaryawanTetapUseCase) {
	karyawanTetapHandler := KaryawanTetapHandler{service}
	r.HandleFunc("/KaryawanTetaps", karyawanTetapHandler.GetAllKaryawanTetaps).Methods(http.MethodGet)
	r.HandleFunc("/KaryawanTetap/{id}", karyawanTetapHandler.KaryawanTetapById).Methods(http.MethodGet)
	r.HandleFunc("/KaryawanTetap", karyawanTetapHandler.CreateKaryawanTetap).Methods(http.MethodPost)
	r.HandleFunc("/KaryawanTetap/{id}", karyawanTetapHandler.DeleteKaryawanTetap).Methods(http.MethodDelete)
	r.HandleFunc("/KaryawanTetap/{id}", karyawanTetapHandler.UpdateKaryawanTetap).Methods(http.MethodPut)
}

func (kt KaryawanTetapHandler) GetAllKaryawanTetaps(w http.ResponseWriter, r *http.Request) {
	KaryawanTetaps, err := kt.KaryawanTetapUseCase.GetAllFixEmployee()
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfKaryawanTetaps, _ := json.Marshal(KaryawanTetaps)
	if err != nil {
		w.Write([]byte("Oops, something when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfKaryawanTetaps)
}

func (kt KaryawanTetapHandler) KaryawanTetapById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	fmt.Println(param["id"])
	KaryawanTetap, err := kt.KaryawanTetapUseCase.GetById(param["id"])
	if err != nil {
		w.Write([]byte("Data not found"))
	}
	byteOfKaryawanTetaps, _ := json.Marshal(KaryawanTetap)
	if err != nil {
		w.Write([]byte("Oops, somethin when wrong !"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfKaryawanTetaps)
}

func (p KaryawanTetapHandler) CreateKaryawanTetap(w http.ResponseWriter, r *http.Request) {
	var KaryawanTetap models.KaryawanTetap
	err := json.NewDecoder(r.Body).Decode(&KaryawanTetap)

	var isTrue = utils.ValidationKaryawanTetap(&KaryawanTetap)
	tools.PrintlnErr(err)
	if isTrue == false {
		w.Write([]byte("Name or address must be fill"))
	} else {
		err = p.KaryawanTetapUseCase.InsertFixEmployee(KaryawanTetap)
		tools.PrintlnErr(err)
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}
}

func (p KaryawanTetapHandler) DeleteKaryawanTetap(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	err := p.KaryawanTetapUseCase.DeleteFixEmployeeById(param["id"])
	tools.PrintlnErr(err)
	log.Println("Delete successful")
	w.Write([]byte("Delete successful"))
}

func (p KaryawanTetapHandler) UpdateKaryawanTetap(w http.ResponseWriter, r *http.Request) {
	var KaryawanTetap models.KaryawanTetap
	param := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&KaryawanTetap)
	var isTrue = utils.ValidationKaryawanTetap(&KaryawanTetap)
	
	if isTrue == false {
		w.Write([]byte("Name or address must be fill"))
	} else {
		err := p.KaryawanTetapUseCase.UpdateFixEmployee(param["id"], KaryawanTetap)
		tools.PrintlnErr(err)
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
