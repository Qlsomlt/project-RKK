package services

import (
	"errors"
	"kode/domain"
	"kode/repositories"
	"kode/utils"
)

type UserService interface {
	Register(name, email, password string) error
	Login(email, password string) (string, error)
	GetAllUsers(requesterRole string) ([]domain.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Register(name, email, password string) error {
	hashed, _ := utils.HashPassword(password)

	user := domain.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     domain.USER,
	}

	return s.repo.Create(&user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID, string(user.Role))
}

func (s *userService) GetAllUsers(requesterRole string) ([]domain.User, error) {
	if requesterRole != "admin" && requesterRole != "superadmin" {
		return nil, errors.New("forbidden")
	}
	return s.repo.FindAll()
}
