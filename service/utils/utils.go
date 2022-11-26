package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func ExtractToken(r *http.Request) (int64, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	token, err := strconv.ParseInt(reqToken, 10, 64)

	if err != nil {
		return -1, err
	}

	return token, nil

}
