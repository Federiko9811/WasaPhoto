package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// PROFILE

	// login
	rt.router.POST("/session", rt.doLogin)
	// update username
	rt.router.PUT("/user/:userId/update-username", rt.authWrapper(rt.setMyUserName))
	// get profile page
	rt.router.GET("/user/:userId/profile-page/:username", rt.authWrapper(rt.getUserProfile))
	// search profiles by username
	rt.router.GET("/user/:userId/search/:username", rt.authWrapper(rt.searchUser))

	// SOCIAL

	// follow user
	rt.router.PUT("/user/:userId/follow/:username", rt.authWrapper(rt.followUser))
	// unfollow user
	rt.router.DELETE("/user/:userId/follow/:username", rt.authWrapper(rt.unfollowUser))
	// ban user
	rt.router.PUT("/user/:userId/ban/:username", rt.authWrapper(rt.banUser))
	// unban user
	rt.router.DELETE("/user/:userId/ban/:username", rt.authWrapper(rt.unbanUser))

	// PHOTO INERACTION

	// return the photos of the followed users
	rt.router.GET("/user/:userId/photos/", rt.authWrapper(rt.getMyStream))
	// post a new photo
	rt.router.POST("/user/:userId/photos/", rt.authWrapper(rt.uploadPhoto))
	// get a photo
	rt.router.GET("/user/:userId/photos/:photoId/", rt.authWrapper(rt.getPhoto))
	// delete a photo
	rt.router.DELETE("/user/:userId/photos/:photoId/", rt.authWrapper(rt.deletePhoto))

	// like a photo
	rt.router.PUT("/user/:userId/photos/:photoId/likes/:authenticatedUserId", rt.authWrapper(rt.likePhoto))
	// unlike a photo
	rt.router.DELETE("/user/:userId/photos/:photoId/likes/:authenticatedUserId", rt.authWrapper(rt.unlikePhoto))
	// get the comments of a photo
	rt.router.GET("/user/:userId/photos/:photoId/comments/", rt.authWrapperNoPath(rt.getPhotoComments))
	// add a comment to a photo
	rt.router.POST("/user/:userId/photos/:photoId/comments/", rt.authWrapperNoPath(rt.commentPhoto))

	// COMMENTS
	rt.router.DELETE("/user/:userId/photos/:photoId/comments/:commentId", rt.authWrapperNoPath(rt.deleteComment))

	return rt.router
}
