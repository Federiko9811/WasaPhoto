package utils

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

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

// ReturnInternalServerError returns an internal server error if an error is detected
func ReturnInternalServerError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
