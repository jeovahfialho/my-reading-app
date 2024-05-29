package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"my-reading-app/internal/domain"
	"my-reading-app/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	Register(input RegisterInput) (string, error)
	Login(input LoginInput) (string, string, error)
	VerifyToken(token string) (*jwt.Token, error)
}

type authService struct {
	userRepo repository.UserRepository
	secret   string
}

func NewAuthService(userRepo repository.UserRepository, secret string) AuthService {
	return &authService{userRepo: userRepo, secret: secret}
}

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *authService) Register(input RegisterInput) (string, error) {
	hashedPassword := hashPassword(input.Password)
	user := domain.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	id, err := a.userRepo.CreateUser(user)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (a *authService) Login(input LoginInput) (string, string, error) {
	user, err := a.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		return "", "", err
	}
	hashedPassword := hashPassword(input.Password)
	if user.Password != hashedPassword {
		return "", "", errors.New("invalid credentials")
	}
	token, err := a.generateToken(user)
	if err != nil {
		return "", "", err
	}
	return token, user.ID, nil
}

func (a *authService) generateToken(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"email":  user.Email,
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secret))
}

func (a *authService) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func hashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
