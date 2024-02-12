package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte("your_secret_key") // Replace with your secret key

func IsAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			tokenString := strings.Split(r.Header["Authorization"][0], " ")[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return MySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
				return
			}

			if token.Valid {
				ctx := context.WithValue(r.Context(), "user", token.Claims)
				handler.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Fprintf(w, "Invalid Authorization Token")
			}
		} else {
			fmt.Fprintf(w, "An Authorization Header is Required")
		}
	})
}
