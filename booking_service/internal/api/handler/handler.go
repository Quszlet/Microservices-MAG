package handler

import (
	"github.com/Quszlet/booking_service/internal/kafka"
	"github.com/Quszlet/booking_service/internal/metrics"
	"github.com/Quszlet/booking_service/internal/service"
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

	booking := r.PathPrefix("/booking").Subrouter()

	booking.HandleFunc("", h.BookTimeSlot).Methods("POST").Name("bookTimeSlot")
	booking.HandleFunc("/fields", h.GetBookingFields).Methods("GET").Name("getBookingFields")
	booking.HandleFunc("/{id:[0-9]+}", h.GetBooking).Methods("GET").Name("getBooking")
	booking.HandleFunc("/{id:[0-9]+}", h.CancelBooking).Methods("PUT").Name("cancelBooking")

	return r
}
