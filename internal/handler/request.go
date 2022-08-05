package handler

import (
	"encoding/json"
	"github.com/Nkez/check/internal/models"
	"net/http"
)

func (h *Handler) PostRequest(w http.ResponseWriter, r *http.Request) {
	var input models.Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := h.service.PostRequest(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h Handler) GetRequest(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	id := v.Get("id")
	response, err := h.service.GetRequest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	output, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetStatus()
	output, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
