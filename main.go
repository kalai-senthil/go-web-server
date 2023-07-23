package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Specify port")
		return
	}
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
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(": %s", PORT),
	}
	log.Printf("Server Running on PORT: %s", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
