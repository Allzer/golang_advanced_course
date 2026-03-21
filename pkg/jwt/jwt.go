package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJwt(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	secret, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		fmt.Println(err)
		return  "", err
	}
	return secret, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	toke_n, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil{
		return false, nil
	}
	email := toke_n.Claims.(jwt.MapClaims)["email"]
	return toke_n.Valid, &JWTData{
		Email: email.(string),
	}
}