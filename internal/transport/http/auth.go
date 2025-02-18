package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		authHeadParts := strings.Split(authHeader[0], " ")

		if len(authHeadParts) != 2 || strings.ToLower(authHeadParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if ValidateToken(authHeadParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
	}
}

func ValidateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
