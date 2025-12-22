package handler

import (
	"github.com/Quszlet/schedule_service/internal/kafka"
	"github.com/Quszlet/schedule_service/internal/metrics"
	"github.com/Quszlet/schedule_service/internal/service"
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

	schedule := r.PathPrefix("/schedule").Subrouter()

	schedule.HandleFunc("", h.GetScheduleDoctor).Methods("GET").Name("getScheduleDoctor")
	schedule.HandleFunc("/fields", h.GetScheduleFields).Methods("GET").Name("getScheduleFields")
	// schedule.HandleFunc("", ).Methods("POST")
	// schedule.HandleFunc("/reserved", ).Methods("PATCH")
	// schedule.HandleFunc("/released", ).Methods("PATCH")
	// schedule.HandleFunc("/delete/{id:[0-9]+}", h.DeactivateDoctor).Methods("PATCH")

	return r
}
