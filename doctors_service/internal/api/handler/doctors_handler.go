package handler

import (
	"log/slog"
	"net/http"
)

func (h *Handler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.CreateDoctor: triggered")
}

func (h *Handler) GetDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.GetDoctor: triggered")
}

func (h *Handler) ActivateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.ActivateDoctor: triggered")
}

func (h *Handler) DeactivateDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.DeactivateDoctor: triggered")
}

func (h *Handler) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	slog.Debug("DoctorService.HandlerAPI.DeleteDoctor: triggered")
}
