package auth_test

import (
	"fmt"
	"http_server/internal/auth"
	"http_server/internal/user"
	"testing"
)

type MockUserRepository struct {
}

func (mur *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, auth.UserExistsError
}

func (mur *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "someEmail",
	}, nil
}
func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "newtest@test.com"
	authService := auth.NewService(&MockUserRepository{})
	email, err := authService.Register(auth.UserData{
		Email:    initialEmail,
		Password: "1",
		Name:     "testUser",
	})
	fmt.Println(email, err)
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Email %s isn't match to %s", email, initialEmail)
	}

}
