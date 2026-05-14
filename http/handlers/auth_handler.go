package handlers

import (
	"github.com/k1rend/RAMS/auth"
	"github.com/k1rend/RAMS/http/dto"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Admin can register a new user with the provided details
// @Tags auth
// @Accept json
// @Param body dto.RegisterRequest true "User registration details"
// @Success 201 {object} repo.User
// @Failure 400 {object} map[string]string{"error": "Invalid request"}
// @Failure 500 {object} map[string]string{"error": "Failed to register user"}
// @Router /auth/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	user, err := h.authService.Register(c.Request().Context(),
		req.Username, req.Email, req.Password, req.FirstName, req.LastName, req.DepartmentId, req.Roles)

	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to register user"})
	}
	return c.JSON(201, user)
}

// Login godoc
// @Summary Login a user
// @Description User can login with username and password to receive access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param body dto.LoginRequest true "User login details"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} map[string]string{"error": "Invalid request"}
// @Failure 401 {object} map[string]string{"error": "Invalid username or password"}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	accessToken, refreshToken, err := h.authService.Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return c.JSON(401, map[string]string{"error": "Invalid username or password"})
	}
	return c.JSON(200, dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

// Refresh godoc
// @Summary Refresh access token
// @Description User can refresh their access token using a valid refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param body dto.RefreshRequest true "Refresh token details"
// @Success 200 {object} dto.RefreshResponse
// @Failure 400 {object} map[string]string{"error": "Invalid request"}
// @Failure 401 {object} map[string]string{"error": "Invalid refresh token"}
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c echo.Context) error {
	var req dto.RefreshRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	accessToken, refreshToken, err := h.authService.Refresh(c.Request().Context(), req.RefreshToken)
	if err != nil {
		return c.JSON(401, map[string]string{"error": "Invalid refresh token"})
	}
	return c.JSON(200, dto.RefreshResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

// TODO: Получение данных о владельцах ресурсов и согласующих
