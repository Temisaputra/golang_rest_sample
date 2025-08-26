package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Temisaputra/warOnk/delivery/repository"
	"github.com/Temisaputra/warOnk/infrastructure/config"
	"github.com/Temisaputra/warOnk/infrastructure/logger"
	"github.com/Temisaputra/warOnk/internal/entity"
	"github.com/Temisaputra/warOnk/pkg/helper"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("token invalid")
	ErrUnauthorized = errors.New("unauthorized")
)

type JwtService interface {
	GenerateToken(user *entity.Users) (string, error)
	ValidateCurrentUser(r *http.Request) (*entity.Users, error)
}

type jwtService struct {
	cfg      config.Config
	log      logger.Logger
	userRepo repository.UserRepository
}

// context key type biar aman tidak tabrakan
type UserContextKey string
type JWTContextKey string

const (
	UserContext UserContextKey = "USER_CONTEXT_KEY"
	JWTContext  JWTContextKey  = "JWT_CONTEXT_KEY"
)

func NewJwtService(cfg config.Config, log logger.Logger, userRepo repository.UserRepository) JwtService {
	return &jwtService{
		cfg:      cfg,
		log:      log,
		userRepo: userRepo,
	}
}

// Claims custom untuk JWT
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken bikin token baru untuk user
func (s *jwtService) GenerateToken(user *entity.Users) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 1) // expired 1 jam
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "warOnk-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

// ValidateCurrentUser validasi JWT dari header Authorization
func (s *jwtService) ValidateCurrentUser(r *http.Request) (*entity.Users, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return nil, helper.NewErrUnauthorized("missing token")
	}

	// format biasanya "Bearer <token>"
	parts := strings.Split(tokenHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil, helper.NewErrUnauthorized("invalid token format")
	}
	tokenString := parts[1]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, helper.NewErrUnauthorized("invalid or expired token")
	}

	// cek expiry manual (opsional karena jwt sudah cek)
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, helper.NewErrUnauthorized("token expired")
	}

	// optional: kalau mau, bisa query ke DB untuk pastikan user masih ada
	// user, err := s.userRepo.FindByID(claims.UserID)
	// if err != nil { return nil, helper.NewErrUnauthorized("user not found") }

	user := &entity.Users{
		ID:    claims.UserID,
		Email: claims.Email,
		Role:  claims.Role,
	}

	return user, nil
}

// SetUserContext masukkan user & jwt token ke dalam context
func SetUserContext(r *http.Request, user *entity.Users) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, UserContext, user)
	ctx = context.WithValue(ctx, JWTContext, r.Header.Get("Authorization"))
	return r.WithContext(ctx)
}

func GetUserContext(ctx context.Context) *entity.Users {
	user, ok := ctx.Value(UserContext).(*entity.Users)
	if !ok {
		return nil
	}
	return user
}

func GetJWTContext(ctx context.Context) string {
	token, _ := ctx.Value(JWTContext).(string)
	return token
}
