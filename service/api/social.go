package api

import (
	"WasaPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) followUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	// Get username from path
	username := p.ByName("username")

	// Check if the username respects the regex
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		utils.ReturnBadRequestMessage(w, nil)
		return
	}

	// Get the token of the username
	token2, err := rt.db.GetUserTokenOnly(username)
	if err != nil {
		utils.ReturnNotFoundError(w)
		return
	}

	// Check if the user is trying to follow himself
	if token == token2 {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Check if the user is banned from the user
	var ban bool
	ban, err = rt.db.CheckBan(token2, token)
	if err != nil || ban {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Check if the user is already following the user
	var check bool
	check, err = rt.db.CheckFollow(token, token2)
	if err != nil || check {
		utils.ReturnConflictMessage(w)
		return
	}

	// Add the follow
	err = rt.db.AddFollow(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	utils.ReturnCreatedMessage(w)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	// Get username from path
	username := p.ByName("username")

	// Check if the username respects the regex
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		utils.ReturnBadRequestMessage(w, nil)
		return
	}

	// Get the token of the username
	token2, err := rt.db.GetUserTokenOnly(username)
	if err != nil {
		utils.ReturnNotFoundError(w)
		return
	}

	// Check if the user is trying to unfollow himself
	if token == token2 {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Check if the user is banned from the user
	var ban bool
	ban, err = rt.db.CheckBan(token2, token)
	if err != nil || ban {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Check if the user is already following the user
	var check bool
	check, err = rt.db.CheckFollow(token, token2)
	if err != nil || !check {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Remove the follow
	err = rt.db.RemoveFollow(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) banUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	// Get username from path
	username := p.ByName("username")

	// Check if the username respects the regex
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		utils.ReturnBadRequestMessage(w, nil)
		return
	}

	// Get the token of the username
	token2, err := rt.db.GetUserTokenOnly(username)
	if err != nil {
		utils.ReturnNotFoundError(w)
		return
	}

	// Check if the user is trying to ban himself
	if token == token2 {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Check if the user is already banned
	var check bool
	check, err = rt.db.CheckBan(token, token2)
	if err != nil || check {
		utils.ReturnConflictMessage(w)
		return
	}

	// Add the ban
	err = rt.db.AddBan(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	utils.ReturnCreatedMessage(w)
}

func (rt *_router) unbanUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	// Get username from path
	username := p.ByName("username")

	// Check if the username respects the regex
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		utils.ReturnBadRequestMessage(w, nil)
		return
	}

	// Get the token of the username
	token2, err := rt.db.GetUserTokenOnly(username)
	if err != nil {
		utils.ReturnNotFoundError(w)
		return
	}

	// Check if the user is trying to unban himself
	if token == token2 {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Check if the user is already banned
	var check bool
	check, err = rt.db.CheckBan(token, token2)
	if err != nil || !check {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Remove the ban
	err = rt.db.RemoveBan(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
