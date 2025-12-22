package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Quszlet/booking_service/internal/metrics"
	"github.com/Quszlet/booking_service/internal/models"
	JSON "github.com/Quszlet/libs/json"
	"github.com/gorilla/mux"
)

func (h *Handler) BookTimeSlot(w http.ResponseWriter, r *http.Request) {
	slog.Debug("BookingService.HandlerAPI.BookTimeSlot: triggered")

	rs := models.RequestBooking{}

	err := JSON.Parse(r, &rs)
	if err != nil {
		metrics.CountHandler("BookTimeSlot", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	err = rs.ValidateCreate()
	if err != nil {
		metrics.CountHandler("BookTimeSlot", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed validate JSON")
		return
	}

	booking_id, err := h.services.Booking.Create(rs)
	if err != nil {
		metrics.CountHandler("BookTimeSlot", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed find schedule")
		return
	}

	metrics.CountHandler("BookTimeSlot", "success")
	JSON.Response(w, http.StatusOK, map[string]int{"booking_id": booking_id})
}

func (h *Handler) GetBooking(w http.ResponseWriter, r *http.Request) {
	slog.Debug("BookingService.HandlerAPI.GetBooking: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	booking, err := h.services.Booking.GetBookingById(id)
	if err != nil { // по timeSlotId
		metrics.CountHandler("GetBooking", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed find booking")
		return
	}

	metrics.CountHandler("GetBooking", "success")
	JSON.Response(w, http.StatusOK, booking)
}

func (h *Handler) GetBookingFields(w http.ResponseWriter, r *http.Request) {
	slog.Debug("BookingService.HandlerAPI.GetBookingFields: triggered")

	bf := models.BookingFields{}

	err := JSON.Parse(r, &bf)
	if err != nil {
		metrics.CountHandler("GetBookingFields", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	if err = bf.Validate(); err != nil {
		metrics.CountHandler("GetBookingFields", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed validate JSON")
		return
	}

	booking, err := h.services.Booking.GetBookingFields(bf.TimeSlotID, bf.Fields)
	if err != nil {
		metrics.CountHandler("GetBookingFields", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed get booking")
		return
	}

	metrics.CountHandler("GetBookingFields", "success")
	JSON.Response(w, http.StatusOK, booking)
}

func (h *Handler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	slog.Debug("BookingService.HandlerAPI.CancelBooking: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if _, err := h.services.Booking.GetBookingById(id); err != nil { // по timeSlotId
		metrics.CountHandler("CancelBooking", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed find booking")
		return
	}

	if err := h.services.Booking.CancelledBookingByTimeSlotId(id); err != nil {
		metrics.CountHandler("CancelBooking", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed delete booking")
		return
	}

	metrics.CountHandler("CancelBooking", "success")
	JSON.Response(w, http.StatusOK, map[string]string{"message": "booking successefull canceled"})
}
