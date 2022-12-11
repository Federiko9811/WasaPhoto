package api

import (
	"WasaPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photo, err := io.ReadAll(r.Body)
	if err != nil || len(photo) == 0 {
		utils.ReturnBadRequestMessage(w, err)
		return
	}
	err = rt.db.PostPhoto(photo, token)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	utils.ReturnCreatedMessage(w)
	return
}

func (rt *_router) deletePhoto(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	// Convert the photo id to int64
	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the request user is the owner of the photo
	check, err := rt.db.CheckPhotoOwner(token, photoId)
	if err != nil || !check {
		utils.ReturnForbiddenMessage(w)
		return
	}

	// Delete the photo
	err = rt.db.DeletePhoto(token, photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (rt *_router) getPhoto(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {

	w.Header().Set("Content-Type", "application/json")

	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}
	var owner int64
	owner, err = rt.db.GetPhotoOwner(photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	ban, err := rt.db.CheckBan(owner, token)
	if err != nil || ban {
		utils.ReturnForbiddenMessage(w)
		return
	}

	var photo []byte
	photo, err = rt.db.GetImage(photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	_, _ = w.Write(photo)
	return
}

func (rt *_router) likePhoto(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var owner int64
	owner, err = rt.db.GetPhotoOwner(photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	var ban bool
	ban, err = rt.db.CheckBan(owner, token)
	if err != nil || ban {
		utils.ReturnForbiddenMessage(w)
		return
	}

	var like bool
	like, err = rt.db.CheckLike(token, photoId)
	if err != nil || like {
		utils.ReturnConfilictMessage(w)
		return
	}

	err = rt.db.LikePhoto(token, photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	utils.ReturnCreatedMessage(w)
	return
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var owner int64
	owner, err = rt.db.GetPhotoOwner(photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	var ban bool
	ban, err = rt.db.CheckBan(owner, token)
	if err != nil || ban {
		utils.ReturnForbiddenMessage(w)
		return
	}

	var like bool
	like, err = rt.db.CheckLike(token, photoId)
	if err != nil || !like {
		utils.ReturnConfilictMessage(w)
		return
	}

	err = rt.db.UnlikePhoto(token, photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
