package api

import (
	"WasaPhoto/service/structs"
	"WasaPhoto/service/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	// Decode request body
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username)

	if err != nil {
		rt.baseLogger.Errorf("Decoding Error: %v", err)
		return
	}
	var identifier structs.Token
	identifier.Identifier, err = rt.db.GetUserToken(username.Username)
	if err != nil {
		rt.baseLogger.Errorf("Error getting identifier: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(identifier)
	if err != nil {
		rt.baseLogger.Errorf("Encoding Error: %v", err)
		return
	}
}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	// Decode request body
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		rt.baseLogger.Errorf("Decoding Error: %v", err)
		return
	}

	// check id username respect a regex
	match, err := regexp.Match(`^[a-zA-Z0-9_-]{3,16}$`, []byte(username.Username))
	if err != nil || !match {
		w.WriteHeader(http.StatusBadRequest)
		res := structs.Message{
			Message: "Error matching regex: %v",
		}
		err = json.NewEncoder(w).Encode(res)
		utils.ReturnInternalServerError(w, err)
		return
	}

	// Get token from header
	token, err := utils.ExtractToken(r)

	// Check if token is valid
	if !rt.db.CheckToken(token) || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		rt.baseLogger.Errorf("Not Valid Token: %v", err)
		res := structs.Message{
			Message: "Not Valid Token",
		}
		err = json.NewEncoder(w).Encode(res)
		utils.ReturnInternalServerError(w, err)
		return
	}

	// Update username in database for the user with the given token
	err = rt.db.SetUserName(token, username.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rt.baseLogger.Errorf("Error setting username: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	res := structs.Message{
		Message: "Username updated",
	}
	err = json.NewEncoder(w).Encode(res)
	utils.ReturnInternalServerError(w, err)
	return
}
