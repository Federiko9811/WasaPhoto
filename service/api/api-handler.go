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

	return rt.router
}
