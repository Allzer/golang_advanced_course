package auth

import (
	"errors"
	"http-server/internal/users"
)

type AuthService struct {
	UserRepository *users.UserRepository
}

func NewAuthService(userRepository *users.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email string, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	user := &users.User{
		Email:    email,
		Password: "",
		Name:     name,
	}
	_, err := service.UserRepository.CreateUser(user)
	if err != nil{
		return "", err
	}
	return user.Email, nil
}
