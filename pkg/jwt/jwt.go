package jwt

import "github.com/golang-jwt/jwt/v5"

type Data struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data Data) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	signedString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", nil
	}
	return signedString, nil
}

func (j *JWT) Parse(tokenString string) (bool, *Data) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}

	email := token.Claims.(jwt.MapClaims)["email"].(string)
	return token.Valid, &Data{
		Email: email,
	}
}
