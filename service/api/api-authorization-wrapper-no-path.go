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

		fn(w, r, ps, token)

	}
}
