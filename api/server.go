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
	router.Handle("/v1/auth/login", middleware.CorsMiddleware(handlers.Login(logger, ss))).Methods("POST")

	api := router.PathPrefix("/v1").Subrouter()
	api.Use(checkMiddleware.GetCheckAuth)
	api.Handle("/profile", middleware.CorsMiddleware(handlers.GetProfile(logger, &settings, ss))).Methods("GET")

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
		Handler: router,
	}
}
