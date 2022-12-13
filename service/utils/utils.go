package utils

import (
	"WasaPhoto/service/structs"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// MESSAGES

// ReturnInternalServerError returns an internal server error if an error is detected
func ReturnInternalServerError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := structs.Message{
			Message: "Internal Server Error " + err.Error(),
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
		return
	}
}

func ReturnCreatedMessage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	res := structs.Message{
		Message: "Created Successfully",
	}
	err := json.NewEncoder(w).Encode(res)
	ReturnInternalServerError(w, err)
}

func ReturnBadRequestMessage(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := structs.Message{
			Message: "Bad Request",
		}
		err = json.NewEncoder(w).Encode(res)
		ReturnInternalServerError(w, err)
		return
	}
}

func ReturnForbiddenMessage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	res := structs.Message{
		Message: "Forbidden Action",
	}
	err := json.NewEncoder(w).Encode(res)
	ReturnInternalServerError(w, err)
}
func ReturnConfilictMessage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusConflict)
	res := structs.Message{
		Message: "This resource is already in the db",
	}
	err := json.NewEncoder(w).Encode(res)
	ReturnInternalServerError(w, err)
}

// FUNCTIONS

func ExtractToken(r *http.Request) (int64, error) {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return -1, errors.New("no token found")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return -1, errors.New("no token found")
	}
	reqToken = splitToken[1]

	token, err := strconv.ParseInt(reqToken, 10, 64)

	if err != nil {
		return -1, err
	}

	return token, nil

}

func ExtractTokenFromPath(w http.ResponseWriter, err error, ps httprouter.Params, paramsName string) int64 {
	pathToken, err := strconv.ParseInt(ps.ByName(paramsName), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		res := structs.Message{
			Message: "Not Valid Token in the Path",
		}
		err = json.NewEncoder(w).Encode(res)
		ReturnInternalServerError(w, err)
		return -1
	}
	return pathToken
}

func CheckUsernameRegex(w http.ResponseWriter, username string) bool {
	match, err := regexp.Match(`^[a-zA-Z0-9_-]{3,16}$`, []byte(username))
	if err != nil || !match {
		w.WriteHeader(http.StatusBadRequest)
		res := structs.Message{
			Message: "Error matching Username regex",
		}
		err = json.NewEncoder(w).Encode(res)
		ReturnInternalServerError(w, err)
		return false
	}
	return true
}
