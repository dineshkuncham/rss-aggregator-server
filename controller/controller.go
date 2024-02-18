package controller

import "github.com/dineshkuncham/rssaggregator/internal/database"

type apiConfig struct {
	DB *database.Queries
}

func NewHandlers(DB *database.Queries) *apiConfig {
	apiCfg := apiConfig{
		DB: DB,
	}
	return &apiCfg
}
