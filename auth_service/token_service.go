package auth_service

import (
	"encoding/json"
	"fmt"
	"mnp-auth-service/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET"))

func LoginHandler(w http.ResponseWriter, r *http.Request, db UserDB) {
	w.Header().Set("Content-Type", "application/json")

	var u models.LoginReguest
	json.NewDecoder(r.Body).Decode(&u)
	user, status, err := db.GetUser(u.Username)
	if err != nil {
		w.WriteHeader(status)
	}

	tokenString, err := createToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf("No username found")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

}

func Grant(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		fn(w, r)
	}
}

func createToken(user models.User) (string, error) {
	claimsMap := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	for _, claim := range user.Claims {
		claimsMap[claim.Key] = claim.Value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMap)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
