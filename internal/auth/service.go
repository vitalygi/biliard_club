package auth

import (
	"biliard_club/internal/user"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository *user.Repository
}

func NewAuthService(userRepository *user.Repository) *Service {
	return &Service{
		UserRepository: userRepository,
	}
}

func (service *Service) Register(phone, password, name string) (*user.User, error) {
	existedUser, _ := service.UserRepository.GetByPhone(phone)
	if existedUser != nil {
		return nil, errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userEntity := &user.User{
		Phone:    phone,
		Password: string(hashedPassword),
		Name:     name,
	}
	_, err = service.UserRepository.Create(userEntity)
	if err != nil {
		return nil, err
	}
	return userEntity, nil
}

func (service *Service) Login(phone, password string) (*user.User, error) {
	existedUser, _ := service.UserRepository.GetByPhone(phone)
	if existedUser == nil {
		return nil, errors.New(ErrWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New(ErrWrongCredentials)
	}
	return existedUser, nil
}
