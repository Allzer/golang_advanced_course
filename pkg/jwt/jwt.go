package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJwt(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	secret, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		fmt.Println(err)
		return  "", err
	}
	return secret, nil
}