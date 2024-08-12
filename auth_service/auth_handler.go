package auth_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mnpt97/mnp-auth-service/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, tokenService *TokenService, db UserDB) {
	w.Header().Set("Content-Type", "application/json")

	var u models.LoginReguest
	json.NewDecoder(r.Body).Decode(&u)
	user, status, err := db.GetUser(u.Username)
	if err != nil {
		w.WriteHeader(status)
	}

	tokenString, err := tokenService.createToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf("No username found")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, tokenService *TokenService, db UserDB) {

}

func Grant(fn func(http.ResponseWriter, *http.Request), tokenService *TokenService, claimValidator func(jwt.MapClaims) bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		status, err := tokenService.verifyToken(tokenString, claimValidator)
		if err != nil {
			w.WriteHeader(status)
			fmt.Fprint(w, err.Error())
			return
		}
		fn(w, r)
	}
}
