package auth_service

import (
	"errors"
	"net/http"
	"time"

	"github.com/mnpt97/mnp-auth-service/models"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	signingMethod    SigningMethod
	expiresInMinutes int
}

type TokenServiceOpt struct {
	expiresInMinutes int
}

func NewTokenService(signingMethod SigningMethod, opts *TokenServiceOpt) *TokenService {
	if opts != nil {

	}
	return &TokenService{
		signingMethod:    signingMethod,
		expiresInMinutes: 24 * 60,
	}

}

func (ts *TokenService) verifyToken(tokenStr string, validateClaims func(jwt.MapClaims) bool) (int, error) {
	token, err := jwt.Parse(tokenStr, ts.signingMethod.KeyFunc)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if !token.Valid {
		return http.StatusForbidden, errors.New("forbidden")
	}
	claims := token.Claims.(jwt.MapClaims)
	if !validateClaims(claims) {
		return http.StatusForbidden, errors.New("forbidden")
	}

	return http.StatusAccepted, nil
}

func (ts *TokenService) createToken(user models.User) (string, error) {
	claimsMap := jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Minute * time.Duration(ts.expiresInMinutes)).Unix(),
	}
	for _, claim := range user.Claims {
		claimsMap[claim.Key] = claim.Value
	}
	var token *jwt.Token
	switch ts.signingMethod.GetMethod() {
	case MethodRSA:
		token = jwt.NewWithClaims(jwt.SigningMethodRS512, claimsMap)
	}
	tokenString, err := token.SignedString(ts.signingMethod.GetSigningKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
