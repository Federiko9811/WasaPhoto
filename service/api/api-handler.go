package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)

	//PROFILE

	//login
	rt.router.POST("/session", rt.doLogin)
	//update username
	rt.router.PUT("/user/:authenticatedUserId/update-username", rt.authWrapper(rt.setMyUserName))
	//get profile page
	rt.router.GET("/user/:authenticatedUserId/profile-page/:username", rt.authWrapper(rt.getUserProfile))
	//search profiles by username
	rt.router.GET("/user/:authenticatedUserId/search/:username", rt.authWrapper(rt.searchUser))

	// SOCIAL

	//follow user
	rt.router.PUT("/user/:authenticatedUserId/follow/:username", rt.authWrapper(rt.followUser))
	//unfollow user
	rt.router.DELETE("/user/:authenticatedUserId/follow/:username", rt.authWrapper(rt.unfollowUser))
	//ban user
	rt.router.PUT("/user/:authenticatedUserId/ban/:username", rt.authWrapper(rt.banUser))
	//unban user
	rt.router.DELETE("/user/:authenticatedUserId/ban/:username", rt.authWrapper(rt.unbanUser))

	//TODO DA FARE

	// PHOTO INERACTION

	//return the photos of the followed users
	rt.router.GET("/user/:authenticatedUserId/photos/", rt.authWrapper(nil))
	//post a new photo
	rt.router.POST("/user/:authenticatedUserId/photos/", rt.authWrapper(rt.uploadPhoto))
	//get a photo
	rt.router.GET("/user/:authenticatedUserId/photos/:photoId/", rt.authWrapper(rt.getPhoto))
	//delete a photo
	rt.router.DELETE("/user/:authenticatedUserId/photos/:photoId/", rt.authWrapper(rt.deletePhoto))

	// TODO CAPIRE COME AGGIUSTARE IL PATH
	////like a photo
	//rt.router.PUT("/user/:username/photos/:photoIds/likes/:authenticatedUserId", rt.authWrapper(nil))
	////unlike a photo
	//rt.router.DELETE("/user/:username/photos/:photoIds/likes/:authenticatedUserId", rt.authWrapper(nil))
	////get the comments of a photo
	//rt.router.GET("/user/:username/photos/:photoIds/comments/owner/:authenticatedUserId", rt.authWrapper(nil))
	////add a comment to a photo
	//rt.router.POST("/user/:username/photos/:photoIds/comments/owner/:authenticatedUserId", rt.authWrapper(nil))
	//
	////COMMENTS
	//
	////get the comment of a photo
	//rt.router.GET("/user/:username/photos/:photoIds/comments/:commentId", rt.authWrapper(nil))
	////delete a comment
	//rt.router.DELETE("/user/:username/photos/:photoIds/comments/:commentId", rt.authWrapper(nil))

	return rt.router
}
