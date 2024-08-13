package infrastructure

import (
	"time"

	"github.com/Tamiru-Alemnew/task-manager/Domain"
	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
    secretKey string
}

func NewJWTService(secretKey string) domain.JWTService {
    return &jwtService{
        secretKey: secretKey,
    }
}

func (s *jwtService) GenerateToken(userID int, username, role string) (string, error) {
    claims := domain.TokenClaims{
        UserID:  string(userID),
        Username: username,
        Role:     role,
        Exp:      time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ParseToken(tokenString string) (*domain.TokenClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.secretKey), nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*domain.TokenClaims)
    if !ok || !token.Valid {
        return nil, jwt.ErrInvalidKey
    }

    return claims, nil
}
