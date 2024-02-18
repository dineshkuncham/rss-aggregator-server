package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dineshkuncham/rssaggregator/controller"
	"github.com/dineshkuncham/rssaggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not found")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to databse:", err)
	}

	queries := database.New(conn)
	apiCfg := controller.NewHandlers(queries)

	go startScraping(queries, 10, time.Minute)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	v1Router := chi.NewRouter()

	v1Router.Get("/health", controller.HandlerReadiness)
	v1Router.Get("/error", controller.HandlerErr)

	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", apiCfg.AuthMiddleware(apiCfg.HandlerGetUser))

	v1Router.Post("/feeds", apiCfg.AuthMiddleware(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)

	v1Router.Post("/feed_follows", apiCfg.AuthMiddleware(apiCfg.HandlerCreateFeedFollows))
	v1Router.Get("/feed_follows", apiCfg.AuthMiddleware(apiCfg.HandlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.AuthMiddleware(apiCfg.HandlerDeleteFeedFollows))

	v1Router.Get("/posts", apiCfg.AuthMiddleware(apiCfg.HandlerGetPostsForUser))

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
