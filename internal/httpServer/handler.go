package httpServer

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v3/middleware/logger"
	authHttp "realty/internal/auth/delivery/http"
	authRepository "realty/internal/auth/repository"
	authUseCase "realty/internal/auth/usecase"
	"realty/internal/middleware"
	realtyHttp "realty/internal/realty/delivery/http"
	realtyRepository "realty/internal/realty/repository"
	realtyUseCase "realty/internal/realty/usecase"
	"realty/pkg/logger"
	storage "realty/pkg/storage/postgres"
)

func (s *Server) MapHandlers(app *fiber.App, logger *logger.ApiLogger) error {
	db, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}

	authRepo := authRepository.NewPostgresRepository(db)
	realtyRepo := realtyRepository.NewPostgresRepository(db)

	authUC := authUseCase.NewAuthUseCase(authRepo)
	realtyUC := realtyUseCase.NewRealtyUseCase(realtyRepo)

	authHandlers := authHttp.NewAuthHandler(authUC, logger)
	realtyHandlers := realtyHttp.NewRealtyHandler(realtyUC, logger)

	app.Use(serverLogger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{},
		AllowHeaders: []string{},
	}))

	mw := middleware.NewMDWManager(s.cfg, logger)

	authGroup := app.Group("auth")
	authHttp.MapAuthRoutes(authGroup, authHandlers)

	realtyGroup := app.Group("realty")
	realtyGroup.Use(mw.JWTMiddleware())
	realtyHttp.MapRealtyRoutes(realtyGroup, realtyHandlers)
	return nil
}
