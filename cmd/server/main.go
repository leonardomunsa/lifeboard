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
	http.HandleFunc("/games", games.GetGames)

	http.ListenAndServe(":8080", nil)
}
