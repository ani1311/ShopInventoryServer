package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../utils"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			json.NewEncoder(w).Encode(utils.MissingAuthTokenResponse)
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.SecretString), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(utils.MissingAuthTokenResponse)
			return
		}

		ctx := context.WithValue(r.Context(), utils.UserString, tk)
		fmt.Println(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
