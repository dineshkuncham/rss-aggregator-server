package main

import (
	"fmt"
	"net/http"

	"github.com/dineshkuncham/rssaggregator/internal/auth"
	"github.com/dineshkuncham/rssaggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
