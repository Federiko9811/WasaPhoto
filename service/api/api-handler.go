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
	rt.router.GET("/profile-page/:username", rt.getUserProfile)

	return rt.router
}
