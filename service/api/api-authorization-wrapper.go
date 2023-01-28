package api

import (
	"WasaPhoto/service/structs"
	"WasaPhoto/service/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params, int64)

func (rt *_router) authWrapper(fn httpRouterHandler) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("content-type", "application/json")
		// Take the content of the Authorization header
		token, err := utils.ExtractToken(r)

		// Check if token is valid
		if !rt.db.CheckToken(token) || err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			rt.baseLogger.Errorf("Not Valid Token: %v", err)
			res := structs.Message{
				Message: "Not Valid Token",
			}
			err = json.NewEncoder(w).Encode(res)
			utils.ReturnInternalServerError(w, err)
			return
		}

		nomiParametri := [3]string{"", "", ""}

		for i, param := range ps {
			nomiParametri[i] = param.Key
		}

		var pathToken int64
		if nomiParametri[0] == "userId" && nomiParametri[2] != "authenticatedUserId" {
			// Get the user id from the token because the request user id is in the userId parameter
			pathToken = utils.ExtractTokenFromPath(w, ps, "userId")
		} else if nomiParametri[0] == "userId" && nomiParametri[2] == "authenticatedUserId" {
			// Get the user id from the token because the request user id is in the authenticatedUserId parameter
			pathToken = utils.ExtractTokenFromPath(w, ps, "authenticatedUserId")
		}

		if pathToken != token {
			w.WriteHeader(http.StatusForbidden)
			rt.baseLogger.Errorf("The path tokens and the auth token aren't equals: %v", err)
			res := structs.Message{
				Message: "The path tokens and the auth token aren't equals",
			}
			err = json.NewEncoder(w).Encode(res)
			utils.ReturnInternalServerError(w, err)
			return
		}

		fn(w, r, ps, pathToken)

	}
}
