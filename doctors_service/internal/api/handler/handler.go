package handler

import (
	"github.com/Quszlet/doctors_service/internal/kafka"
	"github.com/Quszlet/doctors_service/internal/metrics"
	"github.com/Quszlet/doctors_service/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services  *service.Service
	producers *kafka.Producers
}

func NewHandler(s *service.Service, p *kafka.Producers) *Handler {
	return &Handler{services: s, producers: p}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(metrics.HTTPMetricsMiddleware)
	r.Handle("/metrics", metrics.Handler())

	doctors := r.PathPrefix("/doctor").Subrouter()

	doctors.HandleFunc("", h.CreateDoctor).Methods("POST").Name("createDoctor")
	doctors.HandleFunc("", h.GetDoctorFields).Methods("GET").Name("getDoctorFields")
	doctors.HandleFunc("/{id:[0-9]+}", h.GetDoctor).Methods("GET").Name("getDoctor")
	doctors.HandleFunc("/activate/{id:[0-9]+}", h.ActivateDoctor).Methods("PATCH").Name("activateDoctor")
	doctors.HandleFunc("/deactivate/{id:[0-9]+}", h.DeactivateDoctor).Methods("PATCH").Name("deactivateDoctor")
	doctors.HandleFunc("/delete/{id:[0-9]+}", h.DeleteDoctor).Methods("DELETE").Name("deleteDoctor")

	return r
}
