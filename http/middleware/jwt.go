package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTMiddleware struct {
	secret string
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{
		secret: secret,
	}
}

func (m *JWTMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "1Authorization header missing"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "2Invalid authorization header"})
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(m.secret), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "3Invalid token", "detail": err.Error()})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "4Invalid claims"})
		}

		userIDRaw, ok := claims["sub"]
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "5Invalid user ID"})
		}

		userID, ok := userIDRaw.(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "6Invalid user ID format"})
		}

		c.Set("user_id", userID)
		return next(c)
	}
}