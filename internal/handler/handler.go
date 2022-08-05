package handler

import (
	"github.com/Nkez/check/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/status", h.GetStatus).Methods("GET")
	r.HandleFunc("/id", h.GetRequest).Methods("GET")
	r.HandleFunc("/request", h.PostRequest).Methods("POST")
	return r
}
