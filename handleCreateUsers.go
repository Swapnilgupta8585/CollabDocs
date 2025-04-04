package main

import (
	"encoding/json"

	"net/http"
	"time"

	"github.com/Swapnilgupta8585/CollabDocs/internal/auth"
	"github.com/Swapnilgupta8585/CollabDocs/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
}

func (cfg *apiConfig) handleCreateUsers(w http.ResponseWriter, r *http.Request) {

	// request paramter
	type parameter struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	// response payload
	type response struct {
		User User
	}

	// decode the request body
	reqParam := parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqParam)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error decoding the request Body", err)
		return
	}

	// create user in the database
	user, err := cfg.db.CreateUser(r.Context(), reqParam.Email)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error Creating User in the DB", err)
		return
	}

	// create a hash password for the given password
	hashedPassword, err := auth.HashPassword(reqParam.Password)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error converting the password to a hashed password", err)
		return
	}

	// store the hashed password in the DB
	err = cfg.db.AddHashPassword(r.Context(), database.AddHashPasswordParams{HashedPassword: hashedPassword, ID: user.ID})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error Adding the hash Password to the database", err)
		return
	}

	//respond with user(withour hash password ofcourse)
	RespondWithJSON(w, http.StatusCreated, response{User: User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}})

}
