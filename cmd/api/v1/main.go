// @title Gocionics
// @version 1.0
// @description This is a Socionics task using Golang
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"context"
	"fmt"
	"gocionics/internal/app"
	userH "gocionics/internal/handlers/user"
	userRepo "gocionics/internal/repositories/user"
	userUC "gocionics/internal/usecases/user"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {

	appMain := app.New()

	errGroup, errGroupCtx := errgroup.WithContext(context.Background())

	errGroup.Go(func() error {
		log.Printf("Starting server on :%s", appMain.Config.Port)

		if err := appMain.Server.Serve(); err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}

		return nil
	})

	errGroup.Go(func() error {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(
			sigChan,
			syscall.SIGINT,
		)

		select {
		case sigmar := <-sigChan:
			log.Printf("shutdown signal recieved: %v", sigmar)
		case <-errGroupCtx.Done():
			return errGroupCtx.Err()
		}

		ctxWithTimeout, cancelFn := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancelFn()

		log.Println("initiating graceful shutdown")
		if err := appMain.Server.Stop(ctxWithTimeout); err != nil {
			return fmt.Errorf("error graceful shutdown: %v", err)
		}

		log.Println("Server stopped gracefully")

		return nil
	})

	if err := errGroup.Wait(); err != nil {
		log.Printf("application error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Hello Socionics on Go")

	repo := userRepo.INewUserRepository()
	useCase := userUC.NewUserUseCase(repo)
	controller := userCont.NewUserController(useCase)
	router := userH.POST("/register", controller.Register)
}
