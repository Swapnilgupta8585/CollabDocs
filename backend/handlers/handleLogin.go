package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Swapnilgupta8585/CollabDocs/internal/auth"
	"github.com/Swapnilgupta8585/CollabDocs/internal/database"

)


func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request){

	// request body
	type parameter struct{
		Password string `json:"password"`
		Email string `json:"email"`
	}

	// response payload
	type response struct{
		User User
		Token string `json:"token"`
		RefreshToken string`json:"refresh_token"`

	}

	// decode the request body
	reqParam := parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqParam)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, "Error decoding the request Body", err)
		return
	}

	// get user by email
	user, err := h.Cfg.Db.GetUserByEmail(r.Context(), reqParam.Email)
	if err != nil{
		RespondWithError(w, http.StatusNotFound, "user not found", err)
		return
	}

	// check authentication of the user
	err = auth.CheckHashPassword(reqParam.Password, user.HashedPassword)
	if err != nil{
		RespondWithError(w, http.StatusUnauthorized, "Unauthenticated credentials", err)
		return
	}

	// create accessToken
	accessToken, err := auth.MakeJWT(user.ID, h.Cfg.SecretToken, 1 * time.Hour)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, "Error Creating a JWTtoken", err)
		return
	}

	// create refreshToken
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error Creating a Refresh Token", err)
		return
	}

	// store the refresh token in the database
	refreshTokenFromDB, err := h.Cfg.Db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{RefreshToken: refreshToken, UserID: user.ID, ExpiredAt: time.Now().Add(60*24*time.Hour)})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error Adding Refresh Token in the database", err)
		return
	}
	
	RespondWithJSON(w, http.StatusOK, response{
		User: User{
			ID: user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email: user.Email,
		},
		Token: accessToken,
		RefreshToken: refreshTokenFromDB.RefreshToken,
	})

}