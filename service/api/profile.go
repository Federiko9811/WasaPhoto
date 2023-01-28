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

	// Get the username from the request body
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the username respect a regex
	match := utils.CheckUsernameRegex(w, username.Username)
	if !match {
		return
	}

	// Get user token from database, if the user doesn't exist, create it
	var identifier structs.Token
	identifier.Identifier, err = rt.db.GetUserToken(username.Username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	// Encode response body with the token
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(identifier)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}
}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, _ httprouter.Params, pathToken int64) {
	w.Header().Set("content-type", "application/json")

	// Get the username from the request body
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the username respect the regex
	match := utils.CheckUsernameRegex(w, username.Username)
	if !match {
		return
	}

	// Update username in database for the user with the given token
	err = rt.db.SetUserName(pathToken, username.Username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := structs.Message{
		Message: "Username updated",
	}
	err = json.NewEncoder(w).Encode(res)
	utils.ReturnInternalServerError(w, err)
}

func (rt *_router) getUserProfile(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	// Get the username from the path
	username := p.ByName("username")

	// Check if the username respect the regex
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		return
	}

	// Get user profile from database
	profileToken, err := rt.db.GetUserTokenOnly(username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	// Check if the user is banned
	var banned bool
	banned, err = rt.db.CheckBan(profileToken, token)
	if err != nil || banned {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Get user profile from database
	var profile structs.UserProfile
	profile, err = rt.db.GetUserProfile(username, token)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(profile)
	utils.ReturnInternalServerError(w, err)
}

func (rt *_router) searchUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, _ int64) {
	w.Header().Set("content-type", "application/json")

	// Get the username from the path
	username := p.ByName("username")

	// Check if the username respect the regex
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		return
	}

	// Get users from database the username contains the given username
	users, err := rt.db.GetUsersList(username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	utils.ReturnInternalServerError(w, err)
}
