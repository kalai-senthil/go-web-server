package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
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
	hash := sha256.New()
	hash.Write([]byte("sdvdev"))

	user := database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		ApiKey:    fmt.Sprintf("%x", string(hash.Sum(nil))),
	}
	_, err = dbApi.queries.CreateUser(r.Context(), user)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error %v", err))
	}
	respondWithJSON(w, 201, user)
}
func (dbApi *DbApi) getUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	if apiKey == "" {
		respondWithError(w, 403, "Unauthorized")
	}
	user, err := dbApi.queries.GetUser(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error %v", err))
	}
	respondWithJSON(w, 201, user)
}
func (dbApi *DbApi) createFeedHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	user, err := dbApi.queries.GetUser(r.Context(), apiKey)
	if apiKey == "" || err != nil {
		respondWithError(w, 403, "Unauthorized")
	}
	type parameters struct {
		Name string `json:name`
		Url  string `json:url`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decoder.Decode(&params)
	feed := database.CreateFeedParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	}
	_, err = dbApi.queries.CreateFeed(r.Context(), feed)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error %v", err))
	}
	respondWithJSON(w, 201, feed)
}

func (dbApi *DbApi) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	user, err := dbApi.queries.GetUser(r.Context(), apiKey)
	if apiKey == "" || err != nil {
		respondWithError(w, 403, "Unauthorized")
	}

	feeds, err := dbApi.queries.GetFeeds(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 403, "Error retriving feeds")
	}
	respondWithJSON(w, 200, feeds)
}
func (dbApi *DbApi) feedFollowHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	user, err := dbApi.queries.GetUser(r.Context(), apiKey)
	if apiKey == "" || err != nil {
		respondWithError(w, 403, "Unauthorized")
	}
	type parameters struct {
		FeedId string `feedId`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)
	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	}
	_, err = dbApi.queries.CreateFeedFollow(r.Context(), feedFollow)
	if err != nil {
		respondWithError(w, 403, "Error retriving feeds")
	}
	respondWithJSON(w, 200, feedFollow)
}

func (dbApi *DbApi) getFeedFollowsHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	user, err := dbApi.queries.GetUser(r.Context(), apiKey)
	if apiKey == "" || err != nil {
		respondWithError(w, 403, "Unauthorized")
	}
	feedFollows, err := dbApi.queries.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 403, "Error retriving feeds")
	}
	respondWithJSON(w, 200, feedFollows)
}
func (dbApi *DbApi) feedUnFollowHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	feedFollowId := chi.URLParam(r, "feedFollowId")
	user, err := dbApi.queries.GetUser(r.Context(), apiKey)
	if apiKey == "" || err != nil {
		respondWithError(w, 403, "Unauthorized")
	}
	err = dbApi.queries.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 403, "Error retriving feeds")
	}
	respondWithJSON(w, 200, struct{}{})
}
func (dbApi *DbApi) getPostsForUserHadler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	user, err := dbApi.queries.GetUser(r.Context(), apiKey)
	if apiKey == "" || err != nil {
		respondWithError(w, 403, "Unauthorized")
	}
	posts, err := dbApi.queries.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 403, "Error retriving feeds")
	}
	respondWithJSON(w, 200, posts)
}
