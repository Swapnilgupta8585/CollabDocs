package main

import (
	"net/http"
	"time"

	"github.com/Swapnilgupta8585/CollabDocs/internal/auth"
	"github.com/google/uuid"
)

type Doc struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	Content string `json:"content"`
}


func (cfg *apiConfig)handleCreateDocs(w http.ResponseWriter, r *http.Request){
	// response struct
	type response struct{
		Doc Doc
	}

	// get the header of request
	header := r.Header

	// get the JWTtoken string
	tokenString, err := auth.GetBearerToken(header)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Error getting the token string from header", err)
		return
	}

	// validate the token string and get the user id
	userId, err := auth.ValidateJWT(tokenString, cfg.secretToken)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorised user", err)
		return
	}

	// create the doc
	doc, err := cfg.db.CreateDoc(r.Context(), userId)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, "Error creating a doc in the database", err)
		return
	}

	// respond with the doc
	RespondWithJSON(w, http.StatusCreated, response{
		Doc: Doc{
			ID: doc.ID,
			CreatedAt: doc.CreatedAt,
			UpdatedAt: doc.UpdatedAt,
			UserID: doc.UserID,
		},
	})
}