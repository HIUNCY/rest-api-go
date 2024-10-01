package service

import (
	"errors"

	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.Create(user)
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
