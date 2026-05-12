package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"

	"github.com/k1rend/RAMS/repo"
)

type AuthService struct {
	repo   *repo.Queries
	secret string
}

func NewAuthService(repo *repo.Queries, secret string) *AuthService {
	return &AuthService{repo: repo, secret: secret}
}

func (s *AuthService) Register(
	ctx context.Context,
	username string,
	email string,
	password string,
	firstName string,
	lastName string,
	departmentID int,
	roles []int32,
) (repo.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return repo.User{}, err
	}
	user, err := s.repo.CreateUser(ctx, repo.CreateUserParams{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
		FirstName:    firstName,
		LastName:     lastName,
		DepartmentID: pgtype.Int4{Int32: int32(departmentID), Valid: true},
	})
	if err != nil {
		return repo.User{}, err
	}
	for _, role := range roles {
		err = s.repo.GiveUserRole(ctx, repo.GiveUserRoleParams{
			UserID: user.ID,
			RoleID: role,
		})
		if err != nil {
			return repo.User{}, err
		}
	}
	return user, nil
}

func (s *AuthService) Login(
	ctx context.Context,
	username string,
	password string,
) (string, string, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", err
	}
	accessToken, err := s.GenerateJWT(ctx, user.ID, "access", time.Hour)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.GenerateJWT(ctx, user.ID, "refresh", 120*time.Hour)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (s *AuthService) GenerateJWT(ctx context.Context, userID int32,
	tokenType string, ttl time.Duration) (string, error) {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub":   string(userID),
		"roles": roles,
		"exp":   time.Now().Add(ttl).Unix(),
		"type":  tokenType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (any, error) {
		return []byte(s.secret), nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "refresh" {
		return "", "", err
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", "", err
	}
	// uid, err := strconv.Atoi(userID)
	uid, err := strconv.Atoi(userID)
	if err != nil {
		return "", "", err
	}

	accessToken, err := s.GenerateJWT(ctx, int32(uid), "access", time.Hour)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err := s.GenerateJWT(ctx, int32(uid), "refresh", 120*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}
