package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Quszlet/doctors_service/internal/models"
	JSON "github.com/Quszlet/libs/json"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.CreateDoctor: triggered")

	doctor := models.Doctor{}

	err := JSON.Parse(r, &doctor)
	if err != nil {
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	err = doctor.ValidateCreate()
	if err != nil {
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed validate JSON")
		return
	}

	id, err := h.services.Doctor.Create(doctor)
	if err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed created doctor")
		return
	}

	message := fmt.Sprintf("User created with id %d", id)

	JSON.Response(w, http.StatusCreated, message)
}

func (h *Handler) GetDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.GetDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.services.Doctor.Get(id)
	if err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed get doctor")
		return
	}

	JSON.Response(w, http.StatusCreated, user)
}

func (h *Handler) ActivateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.ActivateDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	doctor := &models.Doctor{}
	doctor.Id = uint(id)

	active := true
	doctor.Status = &active
	if err := h.services.Doctor.Update(*doctor, fmt.Sprintf("id = %d", id)); err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed activate doctor")
		return
	}

	JSON.Response(w, http.StatusOK, "Doctor activated")
}

func (h *Handler) DeactivateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.DeactivateDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	doctor := &models.Doctor{}
	doctor.Id = uint(id)

	active := true
	doctor.Status = &active

	inactive := false
	doctor.Status = &inactive
	if err := h.services.Doctor.Update(*doctor, fmt.Sprintf("id = %d", id)); err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed deactivate doctor")
		return
	}

	JSON.Response(w, http.StatusOK, "Doctor deactivated")
}

func (h *Handler) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.DeleteDoctor: triggered")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.services.Doctor.Delete(id); err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed delete doctor")
		return
	}

	message := fmt.Sprintf("Doctor deleted with id %d", id)
	JSON.Response(w, http.StatusOK, message)
}
