package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kalai-senthil/go-web-server/internal/database"
)

type DbApi struct {
	queries *database.Queries
}

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Specify port")
		return
	}
	dbConnection, err := connectToDB()
	if err != nil {
		log.Fatal("Not able to connect to database")
	}
	queries := database.New(dbConnection)
	db := DbApi{
		queries: queries,
	}
	go startScaraping(queries, 10, time.Minute)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
		ExposedHeaders:   []string{"Link"},
	}))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, 200, struct{}{})
	})
	router.Post("/users", db.createUserHandler)
	router.Get("/feeds", db.getFeedsHandler)
	router.Post("/feeds", db.createFeedHandler)
	router.Post("/user", db.getUserHandler)
	router.Post("/feed/follow", db.feedFollowHandler)
	router.Delete("/feed/unfollow/{feedFollowId}", db.feedUnFollowHandler)
	router.Get("/feeds/follow", db.feedFollowHandler)
	router.Get("/posts", db.getPostsForUserHadler)
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(": %s", PORT),
	}
	log.Printf("Server Running on PORT: %s", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
