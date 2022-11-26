package api

import (
	"WasaPhoto/service/structs"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	// Decode request body
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username)

	if err != nil {
		rt.baseLogger.Errorf("Decoding Error: %v", err)
		return
	} else {
		var identifier structs.Token
		identifier.Identifier, err = rt.db.GetUserToken(username.Username)
		if err != nil {
			rt.baseLogger.Errorf("Error getting identifier: %v", err)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			err = json.NewEncoder(w).Encode(identifier)
			if err != nil {
				rt.baseLogger.Errorf("Encoding Error: %v", err)
				return
			}
		}
	}
}
