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
	rt.router.PUT("/user/:authenticatedUserId/update-username", rt.setMyUserName)

	return rt.router
}
