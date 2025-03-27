package games

import (
	"encoding/json"
	"net/http"

	"github.com/leonardomunsa/lifeboard/internal/models"
	"github.com/leonardomunsa/lifeboard/internal/storage"
)

type GameRequest struct {
	Title    string `json:"title"`
	Platform string `json:"platform"`
	Status   string `json:"status"`
}

func StoreGame(w http.ResponseWriter, r *http.Request) {
	var req GameRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Platform == "" || req.Status == "" {
		http.Error(w, "Every field is required", http.StatusBadRequest)
		return
	}

	game := models.Game{
		Title:    req.Title,
		Platform: req.Platform,
		Status:   req.Status,
	}

	result := storage.DB.Create(&game)
	if result.Error != nil {
		http.Error(w, "Error trying to save the game", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	var games []models.Game
	result := storage.DB.Find(&games)

	if result.Error != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}
