package handler

import (
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
