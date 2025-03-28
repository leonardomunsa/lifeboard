package main

import (
	"fmt"
	"github.com/leonardomunsa/lifeboard/internal/games"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/leonardomunsa/lifeboard/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	storage.InitDatabase()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "LifeBoard API funcionando!")
	})
	http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			games.StoreGame(w, r)
		} else if r.Method == http.MethodGet {
			games.GetGames(w, r)
		} else {
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/games/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			games.UpdateGame(w, r)
		} else if r.Method == http.MethodDelete {
			games.DeleteGame(w, r)
		} else {
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
