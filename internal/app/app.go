package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
	"go.uber.org/zap"
	"gocionics/config"
	"gocionics/internal/db"
	"gocionics/internal/server"
	"log"
)

type App struct {
	Config *config.Config
	Server *server.Server
}

func New(cfg *config.Config, router *gin.Engine) *App {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	pgDB, err := db.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	if err := goose.Up(pgDB.DB, "internal/db/migrations"); err != nil {
		logger.Fatal("failed to apply migrations", zap.Error(err))
	}

	return &App{
		Config: cfg,
		Server: server.New(":"+cfg.Port, router),
	}
}
