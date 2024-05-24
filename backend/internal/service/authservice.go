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
	Register(user domain.User) (string, error)
	Login(email, password string) (string, error)
	VerifyToken(token string) (*jwt.Token, error)
}

type authService struct {
	userRepo repository.UserRepository
	secret   string
}

func NewAuthService(userRepo repository.UserRepository, secret string) AuthService {
	return &authService{userRepo: userRepo, secret: secret}
}

func (a *authService) Register(user domain.User) (string, error) {
	hashedPassword := hashPassword(user.Password)
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	id, err := a.userRepo.CreateUser(user)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (a *authService) Login(email, password string) (string, error) {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	hashedPassword := hashPassword(password)
	if user.Password != hashedPassword {
		return "", errors.New("invalid credentials")
	}
	token, err := a.generateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authService) generateToken(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
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
