package auth

import (
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

func (service *Service) Register(email, password, name string) (string, error) {
	_, err := service.UserRepository.FindByEmail(email)
	if err == nil { // it`s bad
		return "", UserExistsError
	}
	newUser := &user.User{
		Email:    email,
		Password: "password",
		Name:     name,
	}
	_, err = service.UserRepository.Create(newUser)
	if err != nil {
		return "", UserRegistrationError
	}
	return newUser.Email, nil
}
