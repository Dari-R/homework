// pkg/jwt/jwt.go
package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Phone string `json:"phone"`
}

type JWT struct{ Secret string }

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": data.Phone,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		phone := claims["phone"].(string)
		return true, &JWTData{Phone: phone}
	}

	return false, nil
}
