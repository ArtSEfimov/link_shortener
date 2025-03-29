package auth

import (
	"golang.org/x/crypto/bcrypt"
	"http_server/internal/user"
)

type Service struct {
	UserRepository *user.Repository
}

func NewService(userRepository *user.Repository) *Service {
	return &Service{
		UserRepository: userRepository,
	}
}

func (service *Service) Register(userData userData) (string, error) {
	_, err := service.UserRepository.FindByEmail(userData.email)
	if err == nil { // it`s bad
		return "", UserExistsError
	}
	hashedPassword, hashedErr := bcrypt.GenerateFromPassword([]byte(userData.password), bcrypt.DefaultCost)
	if hashedErr != nil {
		return "", hashedErr
	}

	newUser := &user.User{
		Email:    userData.email,
		Password: string(hashedPassword),
		Name:     userData.name,
	}
	_, err = service.UserRepository.Create(newUser)
	if err != nil {
		return "", UserRegistrationError
	}
	return newUser.Email, nil
}

func (service *Service) Login(userData userData) (string, error) {
	existedUser, err := service.UserRepository.FindByEmail(userData.email)
	if err != nil {
		return "", UserLoginError
	}
	compareErr := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(userData.password))
	if compareErr != nil {
		return "", UserLoginError
	}
	return existedUser.Email, nil
}
