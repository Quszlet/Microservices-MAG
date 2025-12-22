package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Quszlet/doctors_service/internal/metrics"
	"github.com/Quszlet/doctors_service/internal/models"
	JSON "github.com/Quszlet/libs/json"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.CreateDoctor: triggered")

	doctor := models.Doctor{}

	err := JSON.Parse(r, &doctor)
	if err != nil {
		metrics.CountHandler("CreateDoctor", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	err = doctor.ValidateCreate()
	if err != nil {
		metrics.CountHandler("CreateDoctor", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed validate JSON")
		return
	}

	id, err := h.services.Doctor.Create(doctor)
	if err != nil {
		metrics.CountHandler("CreateDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed created doctor")
		return
	}

	// Send event to Kafka
	h.producers.DoctorsEvents.SendMessage(r.Context(), "doctor.created", map[string]int{"doctor_id": id})

	message := fmt.Sprintf("User created with id %d", id)

	metrics.CountHandler("CreateDoctor", "success")
	JSON.Response(w, http.StatusCreated, message)
}

func (h *Handler) GetDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.GetDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.services.Doctor.Get(id)
	if err != nil {
		metrics.CountHandler("GetDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed get doctor")
		return
	}

	metrics.CountHandler("GetDoctor", "success")
	JSON.Response(w, http.StatusCreated, user)
}

func (h *Handler) GetDoctorFields(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.GetDoctorFields: triggered")

	df := &models.DoctorFields{}

	err := JSON.Parse(r, &df)
	if err != nil {
		metrics.CountHandler("GetDoctorFields", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	err = df.Validate()
	if err != nil {
		metrics.CountHandler("GetDoctorFields", "error")
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed validate JSON")
		return
	}

	// Сделать обертку в сервисе, которая будет обрабатывать получаемую мапу
	// Если поле нет возвращаем map[string]any для существующих и отдаем в запрос
	doctor, err := h.services.Doctor.GetDoctorFields(df.Id, df.Fields)
	if err != nil {
		metrics.CountHandler("GetDoctorFields", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed get doctor")
		return
	}

	metrics.CountHandler("GetDoctor", "success")
	JSON.Response(w, http.StatusCreated, doctor)
}

func (h *Handler) ActivateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.ActivateDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if _, err := h.services.Doctor.Get(id); err != nil {
		metrics.CountHandler("ActivateDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed find doctor")
		return
	}

	doctor := &models.Doctor{}
	doctor.Id = uint(id)

	active := true
	doctor.Status = &active
	if err := h.services.Doctor.Update(*doctor, fmt.Sprintf("id = %d", id)); err != nil {
		metrics.CountHandler("ActivateDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed activate doctor")
		return
	}

	// Send event to Kafka
	h.producers.DoctorsEvents.SendMessage(r.Context(), "doctor.active_changed", map[string]interface{}{"doctor_id": id, "active": true})

	metrics.CountHandler("ActivateDoctor", "success")
	JSON.Response(w, http.StatusOK, "Doctor activated")
}

func (h *Handler) DeactivateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.DeactivateDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if _, err := h.services.Doctor.Get(id); err != nil {
		metrics.CountHandler("DeactivateDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed find doctor")
		return
	}

	doctor := &models.Doctor{}
	doctor.Id = uint(id)

	active := true
	doctor.Status = &active

	inactive := false
	doctor.Status = &inactive
	if err := h.services.Doctor.Update(*doctor, fmt.Sprintf("id = %d", id)); err != nil {
		metrics.CountHandler("DeactivateDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed deactivate doctor")
		return
	}

	// Send event to Kafka
	h.producers.DoctorsEvents.SendMessage(r.Context(), "doctor.active_changed", map[string]interface{}{"doctor_id": id, "active": false})

	metrics.CountHandler("DeactivateDoctor", "success")
	JSON.Response(w, http.StatusOK, "Doctor deactivated")
}

func (h *Handler) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.DeleteDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.services.Doctor.Delete(id); err != nil {
		metrics.CountHandler("DeleteDoctor", "error")
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed delete doctor")
		return
	}

	h.producers.DoctorsEvents.SendMessage(r.Context(), "doctor.deleted", map[string]int{"doctor_id": id})

	respon := map[string]any{}
	respon["message"] = fmt.Sprintf("Doctor deleted with id %d", id)
	metrics.CountHandler("DeleteDoctor", "success")
	JSON.Response(w, http.StatusOK, respon)
}
