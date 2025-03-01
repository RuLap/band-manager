package handlers

import (
	"band-manager/services/band-service/internal/dto"
	"band-manager/services/band-service/internal/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type BandHandler struct {
	bandService services.BandService
}

func NewBandHandler(bandService services.BandService) *BandHandler {
	return &BandHandler{bandService: bandService}
}

func (h *BandHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *dto.CreateBandDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.bandService.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *BandHandler) GetBand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Неверный идентификатор группы", http.StatusBadRequest)
		return
	}

	user, err := h.bandService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
