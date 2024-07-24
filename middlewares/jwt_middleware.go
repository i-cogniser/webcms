package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"webcms/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(authService services.AuthService) echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ContextKey: "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &jwt.RegisteredClaims{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Printf("JWT Error: %v\n", err) // Добавлено для отладки
			if strings.Contains(err.Error(), "token is expired") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token has expired"})
			}
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		},
	}
	return echojwt.WithConfig(config)
}
