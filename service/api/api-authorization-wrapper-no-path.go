package api

import (
	"WasaPhoto/service/structs"
	"WasaPhoto/service/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) authWrapperNoPath(fn httpRouterHandler) func(http.ResponseWriter, *http.Request, httprouter.Params) {
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

		// Check if token is valid
		if !rt.db.CheckToken(token) {
			w.WriteHeader(http.StatusNotFound)
			rt.baseLogger.Errorf("Not Valid Token: %v", err)
			res := structs.Message{
				Message: "Not Valid Token",
			}
			err = json.NewEncoder(w).Encode(res)
			utils.ReturnInternalServerError(w, err)
			return
		}

		fn(w, r, ps, token)

	}
}
