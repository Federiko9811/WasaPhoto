package api

import (
	"WasaPhoto/service/structs"
	"WasaPhoto/service/utils"
	"encoding/json"
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
}

func (rt *_router) deletePhoto(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	// Convert the photo id to int64
	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the photo exists
	var existence bool
	existence, err = rt.db.CheckPhotoExistence(photoId)
	if err != nil || !existence {
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

	// Check if the photo exists and the owner is correct
	var existence bool
	existence, err = rt.db.CheckPhotoExistence(photoId)
	if err != nil || !existence {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var ban bool
	ban, err = rt.db.CheckBan(owner, token)
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
	_, err = w.Write(photo)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}
}

func (rt *_router) likePhoto(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	pathOwner, err := strconv.ParseInt(p.ByName("userId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the photo exists and the owner is correct
	var existence bool
	existence, err = rt.db.CheckPhotoExistence(photoId)
	if err != nil || !existence {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var owner int64
	owner, err = rt.db.GetPhotoOwner(photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	if owner != pathOwner {
		utils.ReturnBadRequestCustomMessage(w)
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
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	pathOwner, err := strconv.ParseInt(p.ByName("userId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the photo exists and the owner is correct
	var existence bool
	existence, err = rt.db.CheckPhotoExistence(photoId)
	if err != nil || !existence {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var owner int64
	owner, err = rt.db.GetPhotoOwner(photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	if owner != pathOwner {
		utils.ReturnBadRequestCustomMessage(w)
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
}

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the photo exists and the owner is correct
	var existence bool
	existence, err = rt.db.CheckPhotoExistence(photoId)
	if err != nil || !existence {
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

	var comment structs.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var newId int64
	newId, err = rt.db.CommentPhoto(token, photoId, comment.Comment)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := structs.CreatedCommentMessage{
		CommentId: newId,
	}
	err = json.NewEncoder(w).Encode(res)
	utils.ReturnInternalServerError(w, err)
}

func (rt *_router) getPhotoComments(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photoId, err := strconv.ParseInt(p.ByName("photoId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	// Check if the photo exists and the owner is correct
	var existence bool
	existence, err = rt.db.CheckPhotoExistence(photoId)
	if err != nil || !existence {
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

	var comments []structs.FullDataComment
	comments, err = rt.db.GetPhotoComments(photoId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}
}

func (rt *_router) getComment(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	commentId, err := strconv.ParseInt(p.ByName("commentId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var owner int64
	owner, err = rt.db.GetCommentOwner(commentId)
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

	var comment structs.FullDataComment
	comment, err = rt.db.GetComment(commentId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}
}

func (rt *_router) deleteComment(w http.ResponseWriter, _ *http.Request, p httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	commentId, err := strconv.ParseInt(p.ByName("commentId"), 10, 64)
	if err != nil {
		utils.ReturnBadRequestMessage(w, err)
		return
	}

	var owner int64
	owner, err = rt.db.GetCommentOwner(commentId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	if owner != token {
		utils.ReturnForbiddenMessage(w)
		return
	}

	err = rt.db.DeleteComment(commentId)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) getMyStream(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, token int64) {
	w.Header().Set("Content-Type", "application/json")

	photos, err := rt.db.GetMyStream(token)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(photos)
	if err != nil {
		utils.ReturnInternalServerError(w, err)
		return
	}
}
