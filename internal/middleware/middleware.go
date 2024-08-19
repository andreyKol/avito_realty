package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v3"
	"realty/internal/common/config"
	"realty/pkg/logger"
	"strings"
)

type MDWManager struct {
	cfg    *config.Config
	logger *logger.ApiLogger
}

func NewMDWManager(cfg *config.Config, logger *logger.ApiLogger) *MDWManager {
	return &MDWManager{
		cfg:    cfg,
		logger: logger,
	}
}

func (mw *MDWManager) JWTMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
		}

		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[len("Bearer "):]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte("a_longer_secret_key_for_hs512_that_is_at_least_64_bytes_long"), nil
		})
		if err != nil || !token.Valid {
			mw.logger.Errorf("Invalid token", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		userType, ok := claims["user_type"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		c.Locals("userType", userType)
		return c.Next()
	}
}
