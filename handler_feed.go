package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dineshkuncham/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	},
	)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Can't create feed: %v", err))
	}

	respondWithJSON(w, 201, convertDatabaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Can't create feed: %v", err))
	}

	respondWithJSON(w, 200, convertDatabaseFeedsToFeeds(feeds))
}
