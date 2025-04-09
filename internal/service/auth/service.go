package auth

import (
	"biliard_club/domain"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetByPhone(phone string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

type Service struct {
	UserRepository UserRepository
}

func NewAuthService(userRepository UserRepository) *Service {
	return &Service{
		UserRepository: userRepository,
	}
}

func (service *Service) Register(phone, password, name string) (*domain.User, error) {
	if phone == "" || password == "" || name == "" {
		return nil, errors.New("phone, password, and name are required")
	}
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userEntity := &domain.User{
		Phone:    phone,
		Password: string(hashedPassword),
		Name:     name,
	}
	userEntity, err = service.UserRepository.Create(userEntity)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}

func (service *Service) Login(phone, password string) (*domain.User, error) {
	existedUser, err := service.UserRepository.GetByPhone(phone)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrWrongCredentials
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return nil, domain.ErrWrongCredentials
	}
	return existedUser, nil
}
