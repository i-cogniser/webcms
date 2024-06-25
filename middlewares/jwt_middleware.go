package middlewares

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
	"webcms/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware(authService services.AuthService) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		Claims:      &jwt.StandardClaims{},
		SuccessHandler: func(c echo.Context) {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*jwt.StandardClaims)
			c.Set("userID", claims.Subject)
		},
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			if strings.Contains(err.Error(), "token is expired") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token has expired"})
			}
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		},
	})
}
