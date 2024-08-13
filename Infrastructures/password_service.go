package infrastructure

import (
    "golang.org/x/crypto/bcrypt"
    "yourapp/domain"
)

type passwordService struct {
    cost int
}

func NewPasswordService(cost int) domain.PasswordService {
    return &passwordService{
        cost: cost,
    }
}

func (s *passwordService) HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func (s *passwordService) ComparePassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
