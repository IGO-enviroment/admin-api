package main

import (
	"admin-api/config"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	log.Println("Запустил api")
	settings, _ := config.Read()
	logger := log.Default()
	mainCtx, cancelMainCtx := context.WithCancel(context.Background())
	defer cancelMainCtx()
	getContext := func(d time.Duration) (context.Context, context.CancelFunc) {
		return context.WithTimeout(mainCtx, d)
	}

	app := NewApp(mainCtx, settings, logger)
	err := app.InitServices()
	if err != nil {
		logger.Fatalf("Не удалось инициализировать приложение: %v", err)
		return
	}

	app.Start()
	logger.Printf("Сервис запущен на порте %d\n", settings.Port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	err = app.Stop(getContext)
	if err != nil {
		logger.Fatalf("Не удалось остановить приложение: %v", err)
		return
	}
}
