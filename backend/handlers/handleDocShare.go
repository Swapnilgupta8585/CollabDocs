package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Swapnilgupta8585/CollabDocs/internal/auth"
	"github.com/Swapnilgupta8585/CollabDocs/internal/database"
	"github.com/google/uuid"

)

type Link struct {
	Token      string `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DocID      uuid.UUID `json:"doc_id"`
	Permission string `json:"permission"`
	ExpiresAt  time.Time `json:"expires_at"`
}


func (h *Handler) HandleDocShare(w http.ResponseWriter, r *http.Request){

	// request body
	type parameter struct{
		Permission string `json:"permission"`
	}

	// response struct
	type response struct {
		Link Link
	}

	// decode the request body
	reqParam := parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqParam)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error decoding the request body", err)
		return
	}

	// check if the permission is in valid format or not
	if reqParam.Permission != "editable" && reqParam.Permission!= "viewable"{
		RespondWithError(w, http.StatusBadRequest, "Invalid Permission format", nil)
		return
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
	userId, err := auth.ValidateJWT(tokenString, h.Cfg.SecretToken)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorised user", err)
		return
	}

	// get doc by DocID
	doc_id := r.PathValue("DocID")

	//parse the docID to be an UUID
	DocID, err := uuid.Parse(doc_id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error parsing the DocID", err)
		return
	}

	// get the doc by id from the DB
	doc, err := h.Cfg.Db.GetDocByID(r.Context(), DocID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error getting the doc using doc id from the DB", err)
		return
	}

	// check whether the user is the owner for the doc or not
	if doc.UserID != userId {
		RespondWithError(w, http.StatusForbidden, "user is not the owner of this resource", nil)
		return
	}

	// create a random string for the token
	token := make([]byte, 32)
	_, err = rand.Read(token)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, "Error creating a random token", err)
		return
	}

	// encode the random bits to a string using hexadecimal encoding
	linkToken := hex.EncodeToString(token)

	// create link with expiry of 24 hours
	link, err := h.Cfg.Db.CreateLink(r.Context(), database.CreateLinkParams{
		Token: linkToken,
		DocID: doc.ID,
		Permission: reqParam.Permission,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	})
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, "Error creating the link in the Database", err)
		return
	}

	// response with link
	RespondWithJSON(w, http.StatusCreated, response{Link: Link{
		Token: link.Token,
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
		DocID: link.DocID,
		Permission: link.Permission,
		ExpiresAt: link.ExpiresAt,
	}})
	
}