package service

import (
	"errors"

	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *model.User) (*model.User, error)
	Login(username, password string) (*model.User, error)
	GetUserByID(userID uint) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Register function for creating a new user
func (s *userService) Register(user *model.User) (*model.User, error) {
	// Check if username or email already exists
	_, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user.Role == "" {
		user.Role = "nasabah"
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user model
	user.Password = string(hashedPassword)

	// Save the user in the database
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login function to authenticate a user
func (s *userService) Login(email, password string) (*model.User, error) {
	// Get user by username
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// GetUserByID returns a user by ID
func (s *userService) GetUserByID(userID uint) (*model.User, error) {
	return s.userRepo.GetUserByID(userID)
}
