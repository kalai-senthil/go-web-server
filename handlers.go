package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kalai-senthil/go-web-server/internal/database"
)

func (dbApi *DbApi) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error: %v", err))
	}
	user := database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	}
	_, err = dbApi.queries.CreateUser(r.Context(), user)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error %v", err))
	}
	respondWithJSON(w, 201, user)
}
