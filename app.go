package main

import (
	"admin-api/api"
	"admin-api/api/middleware"
	"admin-api/config"
	"admin-api/postgres"
	"admin-api/usecases/students"
	"admin-api/usecases/universities"
	"context"
	"log"
	"net/http"
	"time"
)

type App struct {
	server   *http.Server
	mainCtx  context.Context
	settings config.Settings
	logger   *log.Logger
}

func NewApp(mainCtx context.Context, settings config.Settings, logger *log.Logger) *App {
	return &App{
		mainCtx:  mainCtx,
		settings: settings,
		logger:   logger,
	}
}

func (a *App) InitServices() error {
	pgDb, err := postgres.NewPostgresConnector(a.settings.PgConnString).Open()
	if err != nil {
		log.Fatal(err)
	}

	studentService := students.NewStudentsService(pgDb, a.settings)
	universitiesService := universities.NewUniversitiesService(pgDb, a.settings)
	checkMiddleware := middleware.NewCheckTokenMiddleware(&a.settings)

	a.server = api.NewServer(
		a.mainCtx,
		a.settings,
		a.logger,
		studentService,
		universitiesService,
		checkMiddleware,
	)

	return nil
}

func (a *App) Start() {
	go func() {
		if err := http.ListenAndServe(":9091", middleware.CorsMiddleware(a.server.Handler)); err != nil {
			a.logger.Fatalf("Server didn't start: %v", err)
		}
	}()
}

func (a *App) Stop(getContext func(time.Duration) (context.Context, context.CancelFunc)) error {
	serverCtx, cancelServerCtx := getContext(time.Second * 15)
	defer cancelServerCtx()

	err := a.server.Shutdown(serverCtx)
	if err != nil {
		a.logger.Fatalf("Server didn't stop: %v", err)
		return err
	}

	return nil
}
