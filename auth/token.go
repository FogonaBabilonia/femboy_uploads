package auth

import (
	"fmt"
	"time"

	"github.com/FogonaBabilonia/femboy_uploads/config"
	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	claims := claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Secret_key))
}

func validateToken(token_string string, clm *claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token_string, clm, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Secret_key), nil
	})
}

func AuthWithToken(token_string string) (bool, string) {
	clm := claims{}
	token, err := validateToken(token_string, &clm)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, ""
		}
	}

	if !token.Valid {
		return false, ""
	}

	return true, clm.Username
}

//func refreshToken() {}
