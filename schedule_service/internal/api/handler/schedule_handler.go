package handler

import (
	"log/slog"
	"net/http"

	JSON "github.com/Quszlet/libs/json"
	"github.com/Quszlet/schedule_service/internal/metrics"
	"github.com/Quszlet/schedule_service/internal/models"
)

func (h *Handler) GetScheduleDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("ScheduleService.HandlerAPI.GetScheduleDoctor: triggered")

	rs := models.RequestSchedule{}

	err := JSON.Parse(r, &rs)
	if err != nil {
		metrics.CountHandler("GetScheduleDoctor", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	err = rs.ValidateCreate()
	if err != nil {
		metrics.CountHandler("GetScheduleDoctor", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed validate JSON")
		return
	}

	schedule, err := h.services.Schedule.GetScheduleByDoctorIdAndDate(rs.DoctorID, rs.Date)
	if err != nil {
		metrics.CountHandler("GetScheduleDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed find schedule")
		return
	}

	metrics.CountHandler("GetScheduleDoctor", "success")
	JSON.Response(w, http.StatusOK, schedule)
}

func (h *Handler) GetScheduleFields(w http.ResponseWriter, r *http.Request) {
	slog.Debug("ScheduleService.HandlerAPI.GetScheduleFields: triggered")

	ts := models.TimeSlotFields{}

	err := JSON.Parse(r, &ts)
	if err != nil {
		metrics.CountHandler("GetScheduleFields", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	if err = ts.Validate(); err != nil {
		metrics.CountHandler("GetScheduleFields", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed validate JSON")
		return
	}

	timeSlot, err := h.services.Schedule.GetTimeSlotFields(ts.Id, ts.Fields)
	if err != nil {
		metrics.CountHandler("GetScheduleFields", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed get schedule")
		return
	}

	metrics.CountHandler("GetScheduleFields", "success")
	JSON.Response(w, http.StatusOK, timeSlot)
}
