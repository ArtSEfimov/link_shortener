package jwt_test

import (
	"http_server/pkg/jwt"
	"testing"
)

func TestJWT_Create(t *testing.T) {
	const email = "test@test.com"
	jwtService := jwt.NewJWT("usrXxxLbtQ+l8eVGlUwBcqd9Sm8UuWnEutA4DV/Ebvs=")
	token, err := jwtService.Create(jwt.Data{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is not valid")
	}
	if data.Email != email {
		t.Fatalf("Email %s is not equal to %s", data.Email, email)
	}

}
