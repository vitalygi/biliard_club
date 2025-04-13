package auth

import (
	"biliard_club/domain"
	"biliard_club/domain/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository domain.UserRepository
}

func NewAuthService(userRepository domain.UserRepository) *Service {
	return &Service{
		UserRepository: userRepository,
	}
}

func (service *Service) Register(phone, password, name string) (*models.User, error) {
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userEntity := &models.User{
		Phone:    phone,
		Password: string(hashedPassword),
		Name:     name,
	}
	u, err := service.UserRepository.Create(userEntity)
	if err != nil && !errors.Is(err, domain.ErrConflict) {
		return nil, domain.ErrInternalServer
	}
	return u, err
}

func (service *Service) Login(phone, password string) (*models.User, error) {
	existedUser, err := service.UserRepository.GetByPhone(phone)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrUnauthorized
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return nil, domain.ErrUnauthorized
	}
	return existedUser, nil
}
