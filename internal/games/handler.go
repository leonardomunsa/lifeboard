package games

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var game models.Game
	result := storage.DB.First(&game, id)
	if result.Error != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	var req GameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Platform == "" || req.Status == "" {
		http.Error(w, "Every field is needed", http.StatusBadRequest)
		return
	}

	game.Title = req.Title
	game.Platform = req.Platform
	game.Status = req.Status

	result = storage.DB.Save(&game)
	if result.Error != nil {
		http.Error(w, "Error trying to save the game", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var game models.Game
	result := storage.DB.First(&game, id)
	if result.Error != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	result = storage.DB.Delete(&game)
	if result.Error != nil {
		http.Error(w, "Error trying to delete game", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
