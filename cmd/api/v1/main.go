// @title Gocionics API
// @version 1.0
// @description API for Socionics Personality Typing
// @host localhost:8080
// @BasePath /api/v1
//
//go:generate swag init -g cmd/api/v1/main.go --output ../../docs/swagger --parseDependency --parseInternal
package main

import (
	"context"
	"fmt"
	"gocionics/config"
	_ "gocionics/docs/swagger"
	"gocionics/internal/app"
	"gocionics/internal/server"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg := config.NewConfig()
	// Создаем роутер с Swagger
	router := server.NewRouter()
	// Инициализируем приложение
	appL := app.New(cfg, router)
	errGroup, ctx := errgroup.WithContext(context.Background())

	// Server goroutine
	errGroup.Go(func() error {
		log.Printf("Starting server on :%s", cfg.Port)
		if err := appL.Server.Serve(); err != nil {
			return fmt.Errorf("server error: %w", err)
		}
		return nil
	})

	// Shutdown goroutine
	errGroup.Go(func() error {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(
			sigChan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)

		select {
		case sigmar := <-sigChan:
			log.Printf("Received signal: %v", sigmar)
		case <-ctx.Done():
			return ctx.Err()
		}

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()

		log.Println("Shutting down server...")
		return appL.Server.Stop(shutdownCtx)
	})

	if err := errGroup.Wait(); err != nil {
		log.Printf("Application error: %v", err)
		os.Exit(1)
	}

}
