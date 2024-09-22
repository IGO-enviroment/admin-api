package handlers

import (
	"admin-api/api/authentication/generation"
	"admin-api/config"
	"net/http"
	"strings"
)

func GetTokenClaims(r *http.Request, settings *config.Settings) (generation.AccessTokenClaims, error) {
	tokenStr := r.Header.Get("Authorization")
	parts := strings.Split(tokenStr, " ")
	tokenStr = parts[1]
	token, err := generation.NewJWTToken(tokenStr, settings)
	if err != nil {
		return generation.AccessTokenClaims{}, err
	}
	claims := token.Claims()

	return claims, nil
}
