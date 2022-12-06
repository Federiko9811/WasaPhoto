package api

import (
	"WasaPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) followUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	username := p.ByName("username")
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		return
	}

	//TODO Capire se bisogna controllare anche qui se i due utenti sono già amici o
	// o se l'utente sta cercando di seguire se stesso

	err := rt.db.AddFollow(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	utils.ReturnCreatedMessage(w)
	return
}

func (rt *_router) unfollowUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	username := p.ByName("username")
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		return
	}

	err := rt.db.RemoveFollow(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (rt *_router) banUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	username := p.ByName("username")
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		return
	}

	err := rt.db.AddBan(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	utils.ReturnCreatedMessage(w)
	return
}

func (rt *_router) unbanUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("content-type", "application/json")

	username := p.ByName("username")
	match := utils.CheckUsernameRegex(w, username)
	if !match {
		return
	}

	err := rt.db.RemoveBan(token, username)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
