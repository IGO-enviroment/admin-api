package api

import (
	"admin-api/api/handlers"
	"admin-api/api/middleware"
	"admin-api/config"
	"admin-api/usecases/students"
	"admin-api/usecases/universities"
	"context"
	"fmt"
	"github.com/rs/cors"
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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})
	router := mux.NewRouter().UseEncodedPath()
	router.Use(c.Handler)
	router.Handle("/v1/auth/login", handlers.Login(logger, ss)).Methods("POST")

	api := router.PathPrefix("/v1").Subrouter()
	api.Use(checkMiddleware.GetCheckAuth)
	api.Handle("/profile", handlers.GetProfile(logger, &settings, ss)).Methods("GET")

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
		Handler: router,
	}
}
