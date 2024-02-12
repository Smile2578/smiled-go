package middleware

import (
	"fmt"
	"net/http"

	"yourapp/models" // Update with the path to your models package

	"github.com/dgrijalva/jwt-go"
)

func CheckRole(requiredRole models.Role) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value("user").(*jwt.Token)
			if !ok {
				http.Error(w, "Could not retrieve user from context", http.StatusForbidden)
				return
			}

			claims, ok := user.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Could not retrieve claims from token", http.StatusForbidden)
				return
			}

			role := models.Role(claims["role"].(string))
			if role != requiredRole {
				http.Error(w, fmt.Sprintf("Unauthorized - This endpoint requires %s role", requiredRole), http.StatusUnauthorized)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}
