package auth

import (
	"time"

	"github.com/FogonaBabilonia/femboy_uploads/config"
	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	ID uint
	jwt.RegisteredClaims
}

func GenerateToken(id uint) (string, error) {
	claims := claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Secret_key))
}
