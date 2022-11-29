package api

import (
	"WasaPhoto/service/structs"
	"WasaPhoto/service/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
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

// setProfile update the profile of the user
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, _ httprouter.Params, pathToken int64) {
	w.Header().Set("content-type", "application/json")

	// Decode request body
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		rt.baseLogger.Errorf("Decoding Error: %v", err)
		return
	}

	// check id username respect a regex
	match := utils.CheckUsernameRegex(w, username.Username)
	if !match {
		return
	}

	//Update username in database for the user with the given token
	err = rt.db.SetUserName(pathToken, username.Username)
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

func (rt *_router) getUserProfile(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	username := p.ByName("username")
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		return
	}

	// Get user profile from database
	profile, err := rt.db.GetUserProfile(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rt.baseLogger.Errorf("Error setting username: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(profile)
	utils.ReturnInternalServerError(w, err)
	return
}
