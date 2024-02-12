package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dineshkuncham/rssaggregator/internal/auth"
	"github.com/dineshkuncham/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	},
	)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Can't create user: %v", err))
	}

	respondWithJSON(w, 201, convertDatabaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.FetchApiKey(r.Header)
	if err != nil {
		respondWithErr(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithErr(w, 404, fmt.Sprintf("Can't find the user: %v", err))
		return
	}
	log.Println("Found the api key for the user")
	respondWithJSON(w, 200, convertDatabaseUserToUser(user))
}
