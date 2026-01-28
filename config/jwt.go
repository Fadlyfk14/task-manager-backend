package config

import (
	"os"
	"strings"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

/*
JWT Middleware
Pakai di route yang butuh login
*/
func JWTMiddleware() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_change_me"
	}

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(secret),
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			msg := strings.ToLower(err.Error())

			if strings.Contains(msg, "missing") || strings.Contains(msg, "malformed") {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token tidak ada / format salah",
					"hint":    "Gunakan Authorization: Bearer <token>",
				})
			}

			if strings.Contains(msg, "expired") || strings.Contains(msg, "invalid") {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token tidak valid atau sudah kedaluwarsa",
				})
			}

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
		},
	})
}

/*
Generate JWT saat login
*/
func GenerateToken(userID, username, role string, expiresMinutes int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_change_me"
	}

	if expiresMinutes <= 0 {
		expiresMinutes = 60
	}

	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"role":     role,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Duration(expiresMinutes) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

/*
Ambil claims dari context (opsional)
*/
func GetClaims(c *fiber.Ctx) (jwt.MapClaims, bool) {
	user := c.Locals("user")
	if user == nil {
		return nil, false
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}
