package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)

	//PROFILE
	rt.router.POST("/session", rt.doLogin)
	rt.router.PUT("/user/:authenticatedUserId/update-username", rt.authWrapper(rt.setMyUserName))
	rt.router.GET("/user/:authenticatedUserId/profile-page/:username", rt.authWrapper(rt.getUserProfile))
	rt.router.GET("/user/:authenticatedUserId/search/:username", rt.authWrapper(rt.searchUser))

	// SOCIAL
	rt.router.PUT("/user/:authenticatedUserId/follow/:username", rt.authWrapper(rt.followUser))
	rt.router.DELETE("/user/:authenticatedUserId/follow/:username", rt.authWrapper(rt.unfollowUser))
	rt.router.PUT("/user/:authenticatedUserId/ban/:username", rt.authWrapper(rt.banUser))
	rt.router.DELETE("/user/:authenticatedUserId/ban/:username", rt.authWrapper(rt.unbanUser))

	return rt.router
}
