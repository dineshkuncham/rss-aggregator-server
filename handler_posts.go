package main

import (
	"fmt"
	"net/http"

	"github.com/dineshkuncham/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't get posts %v", err))
	}
	respondWithJSON(w, 200, convertDatabasePostsToPosts(posts))
}
