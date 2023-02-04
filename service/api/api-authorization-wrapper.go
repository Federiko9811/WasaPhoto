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
		if err != nil || token == -1 {
			w.WriteHeader(http.StatusUnauthorized)
			rt.baseLogger.Errorf("No Token: %v", err)
			res := structs.Message{
				Message: "No Token in the Header",
			}
			err = json.NewEncoder(w).Encode(res)
			utils.ReturnInternalServerError(w, err)
			return
		}

		if !rt.db.CheckToken(token) {
			w.WriteHeader(http.StatusNotFound)
			rt.baseLogger.Errorf("Not Active Token: %v", err)
			res := structs.Message{
				Message: "Not Active Token",
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
