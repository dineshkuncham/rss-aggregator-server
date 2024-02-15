package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dineshkuncham/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}
	if len(params.FeedId) == 0 {
		respondWithErr(w, 400, "FeedId shouldn't be empty")
	}
	feedFollows, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	},
	)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Can't create feed follows: %v", err))
	}

	respondWithJSON(w, 201, convertDatabaseFeedFollowToFeedFollow(feedFollows))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Can't get feed follow: %v", err))
	}

	respondWithJSON(w, 200, convertDatabaseFeedFollowsToFeedFollows(feedFollows))
}
