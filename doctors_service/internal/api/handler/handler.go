package handler

import (
	"github.com/Quszlet/doctors_service/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	doctors := r.PathPrefix("/doctor").Subrouter()

	doctors.HandleFunc("", h.CreateDoctor).Methods("POST")
	doctors.HandleFunc("/{id:[0-9]+}", h.GetDoctor).Methods("GET")
	doctors.HandleFunc("/activate/{id:[0-9]+}", h.ActivateDoctor).Methods("PATCH")
	doctors.HandleFunc("/deactivate/{id:[0-9]+}", h.DeactivateDoctor).Methods("PATCH")
	doctors.HandleFunc("/delete/{id:[0-9]+}", h.DeleteDoctor).Methods("DELETE")

	return r
}
