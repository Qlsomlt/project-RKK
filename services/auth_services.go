package services

import (
	"kode/repositories"
	"kode/utils"
	"strings"
)

type AuthService struct {
	Logout(authHeader string) error
}

type authService struct {
    tokenRepo repository.TokenRepository
}

func NewAuthService(tr repository.TokenRepository) AuthService {
    return &authService{tr}
}

func (s *authService) Logout(authHeader string) error {
    split := strings.Split(authHeader, " ")
    tokenString := split[1]

    claims, err := utils.ValidateToken(tokenString)
    if err != nil {
        return err
    }

    return s.tokenRepo.Blacklist(
        tokenString,
        claims.ExpiresAt.Time.Unix(),
    )
}

