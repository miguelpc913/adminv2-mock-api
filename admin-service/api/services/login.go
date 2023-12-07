package services

import (
	dtoLogin "admin-v2/api/dto/login"
	"admin-v2/api/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func (sm *ServiceManager) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := dtoLogin.LoginReq{}
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	fmt.Println(loginReq.UserName)
	if loginReq.UserName != "adminV2@clorian.com" {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Wrong auth email"})
		return
	}
	if loginReq.Password != "redmonkey78" {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Wrong auth password"})
		return
	}
	if loginReq.Module != "adminV2" {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Wrong auth module"})
		return
	}
	if loginReq.GrantType != "password" {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Wrong auth grant"})
		return
	}
	tokenString, err := CreateJWT(loginReq)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Wrong auth"})
		return
	}
	var token = map[string]interface{}{
		"token": tokenString,
	}
	helpers.WriteJSON(w, http.StatusBadRequest, token)
}

func CreateJWT(loginReq dtoLogin.LoginReq) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Minute * 100000).Unix(),
		"id":        66,
	}
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
