package user

import (
	"context"
	"errors"
	"log"
	"time"

	"math/rand"
)

type Service struct {
	userRepo Repository
}

func NewService(userRepo Repository) *Service {
	return &Service{userRepo: userRepo}
}

// CreateUser create and store users with a generated token to be used for authentication
func (s *Service) CreateUser(cxt context.Context, username string) (*string, error) {
	if username == "" {
		return nil, errors.New("please enter a valid username")
	}

	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("user already exist")
	}

	newToken := generateToken()

	newUser := &User{
		Username: username,
		Token:    newToken,
	}

	err = s.userRepo.InsertUser(newUser)
	if err != nil {
		return nil, err
	}

	log.Printf("New user created with this username: %s", username)

	return &newToken, nil
}

// ValidateToken takes a token and validate it against the store
// In the case when the token is valid this method return the associated user
func (s *Service) ValidateToken(token string) (*User, error) {
	user, err := s.userRepo.FindUserByToken(token)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid token")
	}

	return user, nil
}

// generateToken is used to generate random tokens for users authentication
func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	bytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 36)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(b)
}
