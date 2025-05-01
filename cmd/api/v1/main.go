// @title Gocionics API
// @version 1.0
// @description API for Socionics Personality Typing
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"context"
	"fmt"
	"gocionics/config"
	"gocionics/internal/app"
	authcontroller "gocionics/internal/controllers/auth"
	charactercontroller "gocionics/internal/controllers/character"
	usercontroller "gocionics/internal/controllers/user"
	"gocionics/internal/db"
	characterrepo "gocionics/internal/repositories/character"
	userrepo "gocionics/internal/repositories/user"
	"gocionics/internal/server"
	authusecase "gocionics/internal/usecases/auth"
	characterusecase "gocionics/internal/usecases/character"
	userusecase "gocionics/internal/usecases/user"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 5 * time.Second

func main() {
	cfg := config.NewConfig()

	// Initialize database
	pgDB, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pgDB.Close()

	// Initialize repositories
	userRepo := userrepo.NewPostgresRepository(pgDB)
	charRepo := characterrepo.NewPostgresRepository(pgDB)

	// Initialize use cases
	authUC := authusecase.NewAuthUseCase(userRepo)
	userUC := userusecase.NewUserUseCase(userRepo, charRepo)
	charUC := characterusecase.NewCharacterUseCase(charRepo)

	// Initialize controllers
	authController := authcontroller.NewAuthController(authUC)
	userController := usercontroller.NewUserController(userUC)
	charController := charactercontroller.NewCharacterController(charUC)

	// Setup router
	router := server.NewRouter()
	api := router.Group("/api/v1")
	{
		auth.SetupRoutes(api, authController)
		user.SetupRoutes(api, userController)
		character.SetupRoutes(api, charController)
	}

	// Create and run app
	app := app.New(cfg, router)

	errGroup, ctx := errgroup.WithContext(context.Background())

	// Server goroutine
	errGroup.Go(func() error {
		log.Printf("Starting server on :%s", cfg.Port)
		if err := app.Server.Serve(); err != nil {
			return fmt.Errorf("server error: %w", err)
		}
		return nil
	})

	// Shutdown goroutine
	errGroup.Go(func() error {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-sigChan:
			log.Printf("Received signal: %v", sig)
		case <-ctx.Done():
			return ctx.Err()
		}

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		log.Println("Shutting down server...")
		return app.Server.Stop(shutdownCtx)
	})

	if err := errGroup.Wait(); err != nil {
		log.Printf("Application error: %v", err)
		os.Exit(1)
	}
}
