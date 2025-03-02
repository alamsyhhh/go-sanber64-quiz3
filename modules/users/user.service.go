package users

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(username, password string) (*User, error)
	LoginUser(username, password string) (string, error)
	UpdateUser(id int, username string) (*User, error)
	GetMe(id int) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUserByID(id int) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) RegisterUser(username, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username:   username,
		Password:   string(hashedPassword),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		CreatedBy:  username,
		ModifiedBy: username,
	}
	
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) LoginUser(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("username atau password salah")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"username" : user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userService) UpdateUser(id int, username string) (*User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.ModifiedAt = time.Now()
	user.ModifiedBy = username

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}

	updatedUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *userService) GetMe(id int) (*User, error) {
	return s.repo.GetUserByID(id)
}
