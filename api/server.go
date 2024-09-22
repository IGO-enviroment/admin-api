package api

import (
	"admin-api/api/handlers"
	"admin-api/api/middleware"
	"admin-api/config"
	"admin-api/usecases/students"
	"admin-api/usecases/universities"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	server   *http.Server
	settings *config.Settings
}

func NewServer(
	ctx context.Context,
	settings config.Settings,
	logger *log.Logger,
	ss students.Service,
	us universities.Service,
	checkMiddleware *middleware.CheckTokenManagerMiddleware,
) *http.Server {

	router := mux.NewRouter().UseEncodedPath()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Continue with the request
			next.ServeHTTP(w, r)
		})
	})
	router.Handle("/v1/auth/login", handlers.Login(logger, ss)).Methods("POST", "OPTIONS")

	api := router.PathPrefix("v1").Subrouter()
	api.Use(checkMiddleware.GetCheckAuth)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
		Handler: router,
	}
}
