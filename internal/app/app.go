package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
	"go.uber.org/zap"
	"gocionics/config"
	"gocionics/internal/db"
	"gocionics/internal/server"
	"log"
	"time"

	authcontroller "gocionics/internal/controllers/auth"
	charactercontroller "gocionics/internal/controllers/character"
	usercontroller "gocionics/internal/controllers/user"
	characterrepo "gocionics/internal/repositories/character"
	userrepo "gocionics/internal/repositories/user"
	authusecase "gocionics/internal/usecases/auth"
	characterusecase "gocionics/internal/usecases/character"
	userusecase "gocionics/internal/usecases/user"
)

type App struct {
	Config *config.Config
	Server *server.Server
	DB     *db.PostgresDB
}

func New(cfg *config.Config, router *gin.Engine) *App {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	// 1. Инициализация БД
	pgDB, err := db.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := pgDB.DB.PingContext(ctx); err != nil {
		logger.Fatal("Database connection failed", zap.Error(err))
	}

	// Ждем пока таблицы станут доступны
	for i := 0; i < 10; i++ {
		_, err := pgDB.DB.ExecContext(ctx, "SELECT 1 FROM users LIMIT 1")
		if err == nil {
			break
		}
		logger.Info("Waiting for tables to be ready...")
		time.Sleep(1 * time.Second)
	}

	// 2. Миграции
	if err := goose.SetDialect("postgres"); err != nil {
		logger.Fatal("failed to set dialect", zap.Error(err))
	}
	if err := goose.Up(pgDB.DB, "./internal/db/migrations"); err != nil {
		logger.Fatal("failed to apply migrations", zap.Error(err))
	}

	// 3. Инициализация слоёв приложения
	userRepo := userrepo.NewPostgresRepository(pgDB.DB)
	charRepo := characterrepo.NewPostgresRepository(pgDB.DB)

	authUC := authusecase.NewAuthUseCase(userRepo, "uzumumw")
	userUC := userusecase.NewUserUseCase(userRepo, charRepo)
	charUC := characterusecase.NewCharacterUseCase(charRepo)

	authController := authcontroller.NewAuthController(authUC)
	userController := usercontroller.NewUserController(userUC)
	charController := charactercontroller.NewCharacterController(charUC)

	// 4. Настройка роутинга
	server.SetupRoutes(router, authController, userController, charController)

	return &App{
		Config: cfg,
		Server: server.New(":"+cfg.Port, router),
		DB:     pgDB,
	}

}
