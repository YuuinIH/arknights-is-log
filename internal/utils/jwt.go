package utils

import (
	"time"

	"github.com/YuuinIH/is-log/internal/config"
	"github.com/golang-jwt/jwt/v4"
	u "github.com/google/uuid"
)

type Claims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

func GenerateToken(uuid u.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, Claims{
		UUID: uuid.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			Issuer:    "is-log",
		},
	})
	ss, err := token.SignedString(config.PrivateKey)
	return ss, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return "114514", nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
