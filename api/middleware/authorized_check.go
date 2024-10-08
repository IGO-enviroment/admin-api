package middleware

import (
	"admin-api/api/authentication/generation"
	"admin-api/config"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

type CheckTokenManagerMiddleware struct {
	settings *config.Settings
}

func NewCheckTokenMiddleware(settings *config.Settings) *CheckTokenManagerMiddleware {
	return &CheckTokenManagerMiddleware{settings: settings}
}

func (middleware *CheckTokenManagerMiddleware) GetCheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1/auth") {
			next.ServeHTTP(w, r)
			return
		}
		authScopes, ok := r.Context().Value("authorization.Scopes").([]string)
		if !ok {
			log.Println("can't convert authorizationScopes to []string")
			next.ServeHTTP(w, r)
			return
		}

		for _, a := range authScopes {
			if a == "Bearer" {
				if !validToken(middleware.settings.JwtSecret, r.Header.Get("Authorization")) {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

func validToken(secretJWT, tokenString string) bool {
	tokenString = strings.Replace(tokenString, "+", "-", -1)
	tokenString = strings.Replace(tokenString, "/", "_", -1)
	var claims generation.AccessTokenClaims
	_, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), &claims, func(token *jwt.Token) (interface{}, error) {
		secretBase64 := base64.StdEncoding.EncodeToString([]byte(secretJWT))
		return []byte(secretBase64), nil
	}, jwt.WithPaddingAllowed())

	if err != nil {
		log.Println(err)
		return false
	}

	if claims.CreationTimestamp+claims.TTL < time.Now().UTC().Unix() {
		return false
	}

	return true
}
