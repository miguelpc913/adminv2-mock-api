package AdminMiddleware

import (
	"admin-v2/api/helpers"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

// HTTP middleware setting a value on the request context
func CheckJTW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		substringToRemove := "Bearer "
		tokenString := strings.Replace(tokenHeader, substringToRemove, "", -1)
		secret := os.Getenv("JWT_SECRET")
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the key for validation
			return []byte(secret), nil
		})

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			helpers.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Not authenticated"})
			return
		}

	})
}
