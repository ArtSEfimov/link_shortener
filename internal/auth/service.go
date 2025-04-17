package auth

import (
	"golang.org/x/crypto/bcrypt"
	"http_server/internal/user"
	"http_server/pkg/di"
)

type Service struct {
	UserRepository di.UserRepositoryInterface
}

func NewService(userRepository di.UserRepositoryInterface) *Service {
	return &Service{
		UserRepository: userRepository,
	}
}

func (service *Service) Register(userData UserData) (string, error) {
	_, err := service.UserRepository.FindByEmail(userData.Email)
	if err == nil { // it`s bad
		return "", UserExistsError
	}
	hashedPassword, hashedErr := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if hashedErr != nil {
		return "", hashedErr
	}

	newUser := &user.User{
		Email:    userData.Email,
		Password: string(hashedPassword),
		Name:     userData.Name,
	}
	_, err = service.UserRepository.Create(newUser)
	if err != nil {
		return "", UserRegistrationError
	}
	return newUser.Email, nil
}

func (service *Service) Login(userData UserData) (string, error) {
	existedUser, err := service.UserRepository.FindByEmail(userData.Email)
	if err != nil {
		return "", UserLoginError
	}
	compareErr := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(userData.Password))
	if compareErr != nil {
		return "", UserLoginError
	}
	return existedUser.Email, nil
}
