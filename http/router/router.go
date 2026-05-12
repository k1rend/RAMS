package router

import (
	"github.com/k1rend/RAMS/http/handlers"
	jwt "github.com/k1rend/RAMS/http/middleware"
	"github.com/labstack/echo/v4"

)

func SetupRouter(
	e *echo.Echo,
	authHandler *handlers.AuthHandler,
	jwtMiddleware *jwt.JWTMiddleware,
) {
	api := e.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.POST("/refresh", authHandler.Refresh)

	protected := api.Group("/protected")
	protected.Use(jwtMiddleware.Handle)
}
