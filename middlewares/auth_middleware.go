package middlewares

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := authParts[1]

		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			return c.JSON(http.StatusInternalServerError, "JWT secret not configured")
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("user", claims)

		return next(c)
	}
}

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			var e *echo.HTTPError
			switch {
			case errors.As(err, &e):
				return c.JSON(e.Code, e.Message)
			default:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			}
		}
		return nil
	}
}
