// @title RAMS API
// @version 1.0
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/k1rend/RAMS/auth"
	"github.com/k1rend/RAMS/http/handlers"
	jwt "github.com/k1rend/RAMS/http/middleware"
	"github.com/k1rend/RAMS/http/router"
	"github.com/k1rend/RAMS/repo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

func main() {
	connStr := os.Getenv("DB_URL")
	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	repo := repo.New(dbPool)

	secret := os.Getenv("JWT_SECRET")

	authService := auth.NewAuthService(repo, secret)
	authHandler := handlers.NewAuthHandler(authService)
	jwtMiddleware := jwt.NewJWTMiddleware(secret)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{
			echo.GET, 
			echo.POST, 
			echo.PUT,
			echo.PATCH, 
			echo.DELETE,
			echo.OPTIONS,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowCredentials: true,
	}))

	router.SetupRouter(e, authHandler, jwtMiddleware)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}