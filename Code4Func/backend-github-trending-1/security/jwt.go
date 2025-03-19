package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/duynghiadev/backend-github-trending/model"
)

const JWT_KEY = "hhhgfdshgfhsdgfshjgfshjdgf"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
