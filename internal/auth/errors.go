package auth

import "errors"

var UserExistsError = errors.New("user already exists")
var UserRegistrationError = errors.New("user registration error")
